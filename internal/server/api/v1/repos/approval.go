package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	approvalPostPayload struct {
		UserID string `json:"user_id"`
	}

	approvalPatchPayload struct {
		IsApproved bool `json:"is_approved"`
	}
)

func (r *Repo) ListApprovals(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("The deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the deployment.")
		return
	}

	as, err := r.i.ListApprovals(ctx, d)
	if err != nil {
		r.log.Error("failed to list approvals.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list approvals.")
		return
	}

	gb.Response(c, http.StatusOK, as)
}

func (r *Repo) GetApproval(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		aid = c.Param("aid")
	)

	ap, err := r.i.FindApprovalByID(ctx, atoi(aid))
	if ent.IsNotFound(err) {
		r.log.Warn("The approval is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The apporval is not found.")
		return
	} else if err != nil {
		r.log.Error("It has failed to get the approval.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the approval.")
		return
	}

	gb.Response(c, http.StatusOK, ap)
}

func (r *Repo) GetMyApproval(c *gin.Context) {
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
		r.log.Warn("The deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the deployment.")
		return
	}

	a, err := r.i.FindApprovalOfUser(ctx, d, u)
	if ent.IsNotFound(err) {
		r.log.Warn("The approval is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The approval is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the approval.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the approval.")
		return
	}

	gb.Response(c, http.StatusOK, a)
}

func (r *Repo) CreateApproval(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	p := &approvalPostPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		r.log.Warn("failed to bind the payload.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the payload.")
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("The deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the deployment.")
		return
	}

	approver, err := r.i.FindUserByID(ctx, p.UserID)
	if ent.IsNotFound(err) {
		r.log.Warn("The approver is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "The approver is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the approver.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the approver.")
		return
	}

	// TODO: Lock for the approval of user.
	if a, _ := r.i.FindApprovalOfUser(ctx, d, u); a != nil {
		r.log.Warn("It already has same approval.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "It already has same approval.")
		return
	}

	ap, err := r.i.CreateApproval(ctx, &ent.Approval{
		UserID:       approver.ID,
		DeploymentID: d.ID,
	})
	if err != nil {
		r.log.Error("It has failed to request a approval.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to request a approval.")
		return
	}

	// Get the approval with edges
	if ae, _ := r.i.FindApprovalOfUser(ctx, d, u); ae != nil {
		ap = ae
	}

	gb.Response(c, http.StatusCreated, ap)
}

func (r *Repo) UpdateApproval(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number = c.Param("number")
	)

	p := &approvalPatchPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		r.log.Warn("failed to bind the payload.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the payload.")
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if ent.IsNotFound(err) {
		r.log.Warn("The deployment is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The deployment is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the deployment.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the deployment.")
		return
	}

	a, err := r.i.FindApprovalOfUser(ctx, d, u)
	if ent.IsNotFound(err) {
		r.log.Warn("The approval is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The approval is not found.")
		return
	} else if err != nil {
		r.log.Error("failed to get the approval.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the approval.")
		return
	}

	if p.IsApproved != a.IsApproved {
		a.IsApproved = p.IsApproved
		if a, err = r.i.UpdateApproval(ctx, a); err != nil {
			r.log.Error("failed to update the approval.", zap.Error(err))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to update the approval.")
			return
		}

	}

	// Get the approval with edges
	if ae, _ := r.i.FindApprovalOfUser(ctx, d, u); ae != nil {
		a = ae
	}

	gb.Response(c, http.StatusOK, a)
}

func (r *Repo) DeleteApproval(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		aid = c.Param("aid")
	)

	ap, err := r.i.FindApprovalByID(ctx, atoi(aid))
	if ent.IsNotFound(err) {
		r.log.Warn("The approval is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusNotFound, "The apporval is not found.")
		return
	} else if err != nil {
		r.log.Error("It has failed to get the approval.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the approval.")
		return
	}

	if err := r.i.DeleteApproval(ctx, ap); err != nil {
		r.log.Error("It has failed to delete the approval.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to delete the approval.")
		return
	}

	c.Status(http.StatusOK)
}
