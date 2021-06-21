package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	UserMiddleware struct{}
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		v, _ := c.Get(gb.KeyUser)
		u, _ := v.(*ent.User)

		if !u.Admin {
			c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{
				"message": "Only admin can access.",
			})
		}
	}
}
