package auth

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/pkg/resp"
	"AynaAPI/server/service/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetInfo godoc
// @Summary get current user info
// @Description 获取当前用户信息
// @Tags Auth
// @Produce json
// @Success 200 {object} resp.UserPublicInfo
// @Router /api/v2/auth/info [get]
func GetInfo(c *gin.Context) {
	appG := app.AppGin{C: c}
	user, ok := appG.GetUser()
	if !ok {
		appG.MakeEmptyResponse(http.StatusOK, e.API_ERROR_INVALID_TOKEN)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, resp.PublicInfo{
		Username:        user.Username,
		PermissionLevel: auth_service.New().GetPermissionByUser(user),
	})
	return
}
