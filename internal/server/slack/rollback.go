package slack

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nleeper/goment"
	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/callback"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/event"
	"github.com/gitploy-io/gitploy/vo"
)

const (
	blockDeployment  = "block_deployment"
	actionDeployment = "aciton_deployment"
)

type (
	rollbackViewSubmission struct {
		DeploymentID int
		ApproverIDs  []string
	}

	deploymentAggregation struct {
		envName     string
		deployments []*ent.Deployment
	}
)

// handleRollbackCmd handles rollback command: "/gitploy rollback OWNER/REPO".
func (s *Slack) handleRollbackCmd(c *gin.Context) {
	ctx := c.Request.Context()

	cmd, _ := slack.SlashCommandParse(c.Request)

	// user have to be exist if chat user is found.
	cu, err := s.i.FindChatUserByID(ctx, cmd.UserID)
	if ent.IsNotFound(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, "Slack is not connected with Gitploy.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get chat user.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	args := strings.Split(cmd.Text, " ")

	ns, n, err := parseFullName(args[1])
	if err != nil {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("`%s` is invalid repository format.", args[1]))
		c.Status(http.StatusOK)
		return
	}

	r, err := s.i.FindRepoOfUserByNamespaceName(ctx, cu.Edges.User, ns, n)
	if ent.IsNotFound(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("The `%s` repository is not found.", args[1]))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the repo.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	config, err := s.i.GetConfig(ctx, cu.Edges.User, r)
	if vo.IsConfigNotFoundError(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, "The config file is not found")
		c.Status(http.StatusOK)
		return
	} else if vo.IsConfigParseError(err) {
		postResponseMessage(cmd.ChannelID, cmd.ResponseURL, "The config file is invliad format.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the config.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	perms, err := s.i.ListPermsOfRepo(ctx, r, "", 1, 100)
	if err != nil {
		s.log.Error("It has failed to list permissions.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	cb, err := s.i.CreateCallback(ctx, &ent.Callback{
		Type:   callback.TypeRollback,
		RepoID: r.ID,
	})
	if err != nil {
		s.log.Error("It has failed to create a new callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	as := s.getSucceedDeploymentAggregation(ctx, r, config)

	_, err = slack.New(cu.BotToken).
		OpenViewContext(ctx, cmd.TriggerID, buildRollbackView(cb.Hash, as, perms))
	if err != nil {
		s.log.Error("It has failed to open a dialog.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func buildRollbackView(callbackID string, as []*deploymentAggregation, perms []*ent.Perm) slack.ModalViewRequest {
	groups := []*slack.OptionGroupBlockObject{}

	for _, a := range as {
		options := []*slack.OptionBlockObject{}

		for _, d := range a.deployments {
			created, _ := goment.New(d.CreatedAt)

			options = append(options, &slack.OptionBlockObject{
				Text: &slack.TextBlockObject{
					Type: slack.PlainTextType,
					Text: fmt.Sprintf("#%d - %s deployed at %s", d.ID, d.GetShortRef(), created.FromNow()),
				},
				Value: strconv.Itoa(d.ID),
			})
		}

		groups = append(groups, &slack.OptionGroupBlockObject{
			Label: &slack.TextBlockObject{
				Type: slack.PlainTextType,
				Text: string(a.envName),
			},
			Options: options,
		})
	}

	approvers := []*slack.OptionBlockObject{}
	for _, perm := range perms {
		u := perm.Edges.User
		if u == nil {
			continue
		}

		approvers = append(approvers, &slack.OptionBlockObject{
			Text: &slack.TextBlockObject{
				Type: slack.PlainTextType,
				Text: u.Login,
			},
			Value: u.ID,
		})
	}

	return slack.ModalViewRequest{
		Type:       slack.VTModal,
		CallbackID: callbackID,
		Title: &slack.TextBlockObject{
			Type: slack.PlainTextType,
			Text: "Rollback",
		},
		Submit: &slack.TextBlockObject{
			Type: slack.PlainTextType,
			Text: "Submit",
		},
		Close: &slack.TextBlockObject{
			Type: slack.PlainTextType,
			Text: "Close",
		},
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				slack.InputBlock{
					Type:    slack.MBTInput,
					BlockID: blockDeployment,
					Label: &slack.TextBlockObject{
						Type: slack.PlainTextType,
						Text: "Deployments",
					},
					Element: slack.SelectBlockElement{
						Type:     slack.OptTypeStatic,
						ActionID: actionDeployment,
						Placeholder: &slack.TextBlockObject{
							Type: slack.PlainTextType,
							Text: "Select target deployment",
						},
						OptionGroups: groups,
					},
				},
				slack.InputBlock{
					Type:    slack.MBTInput,
					BlockID: blockApprovers,
					Label: &slack.TextBlockObject{
						Type: slack.PlainTextType,
						Text: "Approvers",
					},
					Optional: true,
					Element: slack.SelectBlockElement{
						Type:     slack.MultiOptTypeStatic,
						ActionID: actionApprovers,
						Placeholder: &slack.TextBlockObject{
							Type: slack.PlainTextType,
							Text: "Select approvers",
						},
						Options: approvers,
					},
				},
			},
		},
	}
}

func (s *Slack) getSucceedDeploymentAggregation(ctx context.Context, r *ent.Repo, cf *vo.Config) []*deploymentAggregation {
	a := []*deploymentAggregation{}

	for _, env := range cf.Envs {
		ds, _ := s.i.ListDeploymentsOfRepo(ctx, r, env.Name, string(deployment.StatusSuccess), 1, 5)
		if len(ds) == 0 {
			continue
		}

		a = append(a, &deploymentAggregation{
			envName:     env.Name,
			deployments: ds,
		})
	}

	return a
}

func (s *Slack) interactRollback(c *gin.Context) {
	ctx := c.Request.Context()

	// InteractionCallbackParse always to be parsed successfully because
	// it was called in the Interact method.
	itr, _ := s.InteractionCallbackParse(c.Request)
	cb, _ := s.i.FindCallbackByHash(ctx, itr.View.CallbackID)

	cu, _ := s.i.FindChatUserByID(ctx, itr.User.ID)

	sm := parseRollbackSubmissions(itr)

	d, err := s.i.FindDeploymentByID(ctx, sm.DeploymentID)
	if err != nil {
		s.log.Error("It has failed to find the deployment.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	cf, err := s.i.GetConfig(ctx, cu.Edges.User, cb.Edges.Repo)
	if vo.IsConfigNotFoundError(err) {
		postBotMessage(cu, "The config file is not found.")
		c.Status(http.StatusOK)
		return
	} else if vo.IsConfigParseError(err) {
		postBotMessage(cu, "The config file is invliad format.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the config.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if !cf.HasEnv(d.Env) {
		postBotMessage(cu, "The environment is not defined.")
		c.Status(http.StatusOK)
		return
	}

	env := cf.GetEnv(d.Env)

	next, err := s.i.GetNextDeploymentNumberOfRepo(ctx, cb.Edges.Repo)
	if err != nil {
		s.log.Error("It has failed to get the next deployment number.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	d, err = s.i.Rollback(ctx, cu.Edges.User, cb.Edges.Repo, &ent.Deployment{
		Number: next,
		Type:   deployment.Type(d.Type),
		Ref:    d.Ref,
		Sha:    d.Sha,
		Env:    d.Env,
	}, env)
	if ent.IsConstraintError(err) {
		postBotMessage(cu, "The conflict occurs, please retry.")
		c.Status(http.StatusOK)
		return
	} else if vo.IsUnprocessibleDeploymentError(err) {
		postBotMessage(cu, fmt.Sprintf("It is unprocessible entity. (Discussion <%s|#64>)", linkUnprocessalbeEntity))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to deploy.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if _, err := s.i.CreateEvent(ctx, &ent.Event{
		Kind:         event.KindDeployment,
		Type:         event.TypeCreated,
		DeploymentID: d.ID,
	}); err != nil {
		s.log.Error("It has failed to create the event.", zap.Error(err))
	}

	if env.IsApprovalEabled() {
		for _, id := range sm.ApproverIDs {
			a, err := s.i.CreateApproval(ctx, &ent.Approval{
				UserID:       id,
				DeploymentID: d.ID,
			})
			if err != nil {
				s.log.Error("It has failed to create a new approval.", zap.Error(err))
				continue
			}

			if _, err := s.i.CreateEvent(ctx, &ent.Event{
				Kind:       event.KindDeployment,
				Type:       event.TypeCreated,
				ApprovalID: a.ID,
			}); err != nil {
				s.log.Error("It has failed to create the event.", zap.Error(err))
			}
		}
	}

	c.Status(http.StatusOK)
}

func parseRollbackSubmissions(itr slack.InteractionCallback) *rollbackViewSubmission {
	sm := &rollbackViewSubmission{}

	values := itr.View.State.Values
	if v, ok := values[blockDeployment][actionDeployment]; ok {
		sm.DeploymentID = atoi(v.SelectedOption.Value)
	}

	ids := make([]string, 0)
	if v, ok := values[blockApprovers][actionApprovers]; ok {
		for _, option := range v.SelectedOptions {
			ids = append(ids, option.Value)
		}

		sm.ApproverIDs = ids
	}

	return sm
}
