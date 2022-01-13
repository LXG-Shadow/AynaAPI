package auth

import (
	"AynaAPI/config"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/middleware/jwt"
	"AynaAPI/server/service/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary login
// @Description 登录
// @Tags Auth
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} app.AppJsonResponse "
// @Router /api/v1/auth/login [get]
func Login(c *gin.Context) {
	appG := app.AppGin{C: c}

	username, b1 := c.GetQuery("username")
	password, b2 := c.GetQuery("password")
	if !b1 {
		appG.MakeEmptyResponse(http.StatusOK, e.AUTH_ERROR_REQUIRE_USERNAME)
		return
	}
	if !b2 {
		appG.MakeEmptyResponse(http.StatusOK, e.AUTH_ERROR_REQUIRE_PASSWORD)
		return
	}
	// todo change to service
	if !auth_service.New().Login(auth_service.LoginParam{
		Username: username,
		Password: password,
	}) {
		appG.MakeEmptyResponse(http.StatusOK, e.AUTH_ERROR_U_P_NOT_MATCH)
		return
	}
	token, err := jwt.GenerateToken(username, password)
	if err != nil {
		appG.MakeEmptyResponse(http.StatusInternalServerError, e.ERROR_UNKNOWN)
		return
	}
	appG.SetCookie(config.ServerConfig.JwtTokenName, token, 7*24*3600, false, true)
	appG.MakeResponse(http.StatusOK, e.AUTH_OK, map[string]string{
		"token": token,
	})
}
