package repos

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/notification"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	deploymentPostPayload struct {
		Type string `json:"type"`
		Ref  string `json:"ref"`
		Env  string `json:"env"`
	}

	deploymentPatchPayload struct {
		Status string `json:"status"`
	}
)

func (r *Repo) ListDeployments(c *gin.Context) {
	var (
		env     = c.Query("env")
		status  = c.Query("status")
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)
	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	ctx := c.Request.Context()

	ds, err := r.i.ListDeployments(ctx, re, env, status, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Error("failed to list deployments.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list deployments.")
		return
	}

	gb.Response(c, http.StatusOK, ds)
}

func (r *Repo) GetDeploymentByNumber(c *gin.Context) {
	var (
		number = c.Param("number")
	)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	ctx := c.Request.Context()

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("the deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
	} else if err != nil {
		r.log.Error("failed to find deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "It has failed to find the deployment.")
	}

	gb.Response(c, http.StatusOK, d)
}

func (r *Repo) CreateDeployment(c *gin.Context) {
	ctx := c.Request.Context()

	p := &deploymentPostPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the payload.")
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	cf, err := r.i.GetConfig(ctx, u, re)
	if vo.IsConfigNotFoundError(err) {
		r.log.Warn("failed to get the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to find the configuraton file.")
		return
	} else if vo.IsConfigParseError(err) {
		r.log.Warn("failed to parse the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to parse the configuraton file.")
		return
	} else if err != nil {
		r.log.Error("failed to get the configuration file.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the configuraton file.")
		return
	}

	if !cf.HasEnv(p.Env) {
		r.log.Warn("failed to get the env.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to get the env in the configuration file.")
		return
	}

	// Prepare to deploy:
	// 1) commit SHA.
	// 2) next deployment number.
	var (
		sha    string
		number int
	)
	if sha, err = r.getCommitSha(ctx, u, re, p.Type, p.Ref); err != nil {
		r.log.Error("failed to get the commit sha.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "The Ref is invalid.")
		return
	}

	if number, err = r.i.GetNextDeploymentNumberOfRepo(ctx, re); err != nil {
		r.log.Error("failed to get the next deployment number.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to the next deployment number.")
		return
	}

	d, err := r.i.Deploy(ctx, u, re, &ent.Deployment{
		Number: number,
		Type:   deployment.Type(p.Type),
		Ref:    p.Ref,
		Sha:    sha,
		Env:    p.Env,
	}, cf.GetEnv(p.Env))
	if ent.IsConstraintError(err) {
		r.log.Warn("deployment number conflict.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusConflict, "The conflict occurs, please retry.")
		return
	} else if err != nil {
		r.log.Error("failed to deploy.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to deploy.")
		return
	}

	if err = r.i.Publish(ctx, notification.TypeDeployment, re, d, nil); err != nil {
		r.log.Warn("failed to notify the deployment.", zap.Error(err))
	}

	// Get the deployment with edges.
	if de, _ := r.i.FindDeploymentByID(ctx, d.ID); de != nil {
		d = de
	}

	gb.Response(c, http.StatusCreated, d)
}

func (r *Repo) UpdateDeployment(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	p := &deploymentPatchPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the payload.")
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("the deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
	}

	cf, err := r.i.GetConfig(ctx, u, re)
	if vo.IsConfigNotFoundError(err) {
		r.log.Warn("failed to get the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to find the configuraton file.")
		return
	} else if vo.IsConfigParseError(err) {
		r.log.Warn("failed to parse the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to parse the configuraton file.")
		return
	} else if err != nil {
		r.log.Error("failed to get the configuration file.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the configuraton file.")
		return
	}

	if !cf.HasEnv(d.Env) {
		r.log.Warn("failed to get the env.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to get the env in the configuration file.")
		return
	}

	if p.Status == string(deployment.StatusCreated) && d.Status == deployment.StatusWaiting {
		// Check the deployment is approved:
		// Approved >= Required Approval Count
		if !r.i.IsApproved(ctx, d) {
			r.log.Warn("The deployment is not approved yet.", zap.Int("deployment_id", d.ID))
			gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It is not approved yet.")
			return
		}

		if d, err = r.i.DeployToSCM(ctx, u, re, d, cf.GetEnv(d.Env)); err != nil {
			r.log.Error("failed to deploy to SCM.", zap.Error(err))
			gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to deploy to SCM.")
			return
		}
	}

	// Get the deployment with edges.
	if de, _ := r.i.FindDeploymentByID(ctx, d.ID); de != nil {
		d = de
	}

	gb.Response(c, http.StatusOK, d)
}

func (r *Repo) RollbackDeployment(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("the deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
	}

	cf, err := r.i.GetConfig(ctx, u, re)
	if vo.IsConfigNotFoundError(err) {
		r.log.Warn("failed to get the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to find the configuraton file.")
		return
	} else if vo.IsConfigParseError(err) {
		r.log.Warn("failed to parse the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to parse the configuraton file.")
		return
	} else if err != nil {
		r.log.Error("failed to get the configuration file.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the configuraton file.")
		return
	}

	if !cf.HasEnv(d.Env) {
		r.log.Warn("failed to get the env.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to get the env in the configuration file.")
		return
	}

	// Prepare to rollback:
	// 1) next deployment number.
	var (
		next int
	)

	if next, err = r.i.GetNextDeploymentNumberOfRepo(ctx, re); err != nil {
		r.log.Error("failed to get the next deployment number.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to the next deployment number.")
		return
	}

	d, err = r.i.Rollback(ctx, u, re, &ent.Deployment{
		Number: next,
		Type:   d.Type,
		Ref:    d.Ref,
		Sha:    d.Sha,
		Env:    d.Env,
	}, cf.GetEnv(d.Env))
	if ent.IsConstraintError(err) {
		r.log.Warn("deployment number conflict.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusConflict, "The conflict occurs, please retry.")
		return
	} else if err != nil {
		r.log.Error("failed to deploy.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to deploy.")
		return
	}

	if err = r.i.Publish(ctx, notification.TypeDeployment, re, d, nil); err != nil {
		r.log.Warn("failed to notify the deployment.", zap.Error(err))
	}

	// Get the deployment with edges.
	if de, _ := r.i.FindDeploymentByID(ctx, d.ID); de != nil {
		d = de
	}

	gb.Response(c, http.StatusCreated, d)
}

func (r *Repo) GetConfig(c *gin.Context) {
	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	ctx := c.Request.Context()

	config, err := r.i.GetConfig(ctx, u, re)
	if vo.IsConfigNotFoundError(err) {
		r.log.Warn("failed to find the config file.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "It has failed to find the configuraton file.")
		return
	} else if vo.IsConfigParseError(err) {
		r.log.Warn("failed to parse the config.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It has failed to parse the configuraton file.")
		return
	} else if err != nil {
		r.log.Error("failed to get the config file.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the config file.")
		return
	}

	gb.Response(c, http.StatusOK, config)
}

func (r *Repo) getCommitSha(ctx context.Context, u *ent.User, re *ent.Repo, typ, ref string) (string, error) {
	if typ == "commit" {
		return ref, nil
	} else if typ == "branch" {
		b, err := r.i.GetBranch(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return b.CommitSha, nil
	} else if typ == "tag" {
		t, err := r.i.GetTag(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return t.CommitSha, nil
	}

	return "", fmt.Errorf("Type must be one of commit, branch, tag.")
}
