package auth

import (
	"AynaAPI/config"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/common"
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
// @Success 200 {object} resp.LoginResp
// @Router /api/v2/auth/login [get]
func Login(c *gin.Context) {
	appG := app.AppGin{C: c}

	var loginPara auth_service.LoginParam

	if ok, errs := appG.BindWithParam(&loginPara); !ok {
		common.Logger.WithField("parameter", loginPara).Warn("Login Validation Fail")
		appG.MakeErrorResponse(http.StatusOK, e.API_ERROR_INVALID_PARAMETER, errs.Errors()...)
		return
	}
	if !auth_service.New().Login(loginPara) {
		appG.MakeErrorResponse(http.StatusOK, e.AUTH_ERROR_U_P_NOT_MATCH)
		return
	}
	token, err := jwt.GenerateToken(loginPara.Username, loginPara.Password)
	if err != nil {
		appG.MakeErrorResponse(http.StatusInternalServerError, e.ERROR_UNKNOWN)
		return
	}
	appG.SetCookie(config.ServerConfig.JwtTokenName, token, 7*24*3600, true, true)
	appG.MakeResponse(http.StatusOK, e.AUTH_OK, map[string]string{
		"token": token,
	})
}
