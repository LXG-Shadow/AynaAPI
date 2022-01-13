package perm

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckPermission(level int) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.AppGin{C: c}
		user, ok := appG.GetUser()
		if !ok {
			appG.MakeEmptyResponse(http.StatusUnauthorized, e.API_ERROR_REQUIRE_TOKEN)
			c.Abort()
			return
		}
		if auth_service.New().GetPermissionByUser(user) < level {
			appG.MakeEmptyResponse(http.StatusUnauthorized, e.API_ERROR_PERMISSION_NOT_ALLOWED)
			c.Abort()
			return
		}
		c.Next()
	}
}
