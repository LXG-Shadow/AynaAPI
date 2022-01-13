package auth

import (
	"AynaAPI/config"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Logout godoc
// @Summary logout
// @Description 登出
// @Tags Auth
// @Produce json
// @Success 200 {object} app.AppJsonResponse "
// @Router /api/v2/auth/logout [get]
func Logout(c *gin.Context) {
	appG := app.AppGin{C: c}
	appG.DeleteCookie(config.ServerConfig.JwtTokenName)
	appG.MakeResponse(http.StatusOK, e.API_OK, map[string]string{
		"msg": "logout ok",
	})
}
