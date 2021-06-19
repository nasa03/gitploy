package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanjunlee/gitploy/ent"
	errs "github.com/hanjunlee/gitploy/internal/errors"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"github.com/hanjunlee/gitploy/vo"
	"go.uber.org/zap"
)

func (r *Repo) ListCommits(c *gin.Context) {
	var (
		branch  = c.Query("branch")
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)

	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	commits, err := r.i.ListCommits(ctx, u, repo, branch, atoi(page), atoi(perPage))
	if err != nil {
		r.log.Error("failed to list commits.", zap.String("repo", repo.Name), zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list commits.")
		return
	}

	gb.Response(c, http.StatusOK, commits)
}

func (r *Repo) GetCommit(c *gin.Context) {
	var (
		sha = c.Param("sha")
	)

	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	commit, err := r.i.GetCommit(ctx, u, repo, sha)
	if errs.IsRefNotFoundError(err) {
		r.log.Warn("The commit is not found.", zap.String("repo", repo.Name), zap.String("sha", sha), zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The commit is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the commit.", zap.String("repo", repo.Name), zap.String("sha", sha), zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the commit.")
		return
	}

	gb.Response(c, http.StatusOK, commit)
}

func (r *Repo) ListStatuses(c *gin.Context) {
	var (
		sha = c.Param("sha")
	)

	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	ss, err := r.i.ListCommitStatuses(ctx, u, repo, sha)
	if errs.IsRefNotFoundError(err) {
		r.log.Warn("The commit is not found.", zap.String("repo", repo.Name), zap.String("sha", sha), zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The commit is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the commit status.", zap.String("repo", repo.Name), zap.String("sha", sha), zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the commit.")
		return
	}

	gb.Response(c, http.StatusOK, map[string]interface{}{
		"state":    mergeState(ss),
		"statuses": ss,
	})
}

func mergeState(ss []*vo.Status) string {
	// The state is failure if one of them is failure.
	for _, s := range ss {
		if s.State == vo.StatusStateFailure {
			return string(vo.StatusStateFailure)
		}
	}

	for _, s := range ss {
		if s.State == vo.StatusStatePending {
			return string(vo.StatusStatePending)
		}
	}

	return string(vo.StatusStateSuccess)
}