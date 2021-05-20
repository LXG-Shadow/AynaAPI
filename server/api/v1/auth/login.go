package auth

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/middleware/jwt"
	"AynaAPI/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	if ok, _ := models.AuthUser(username, password); !ok {
		appG.MakeEmptyResponse(http.StatusOK, e.AUTH_ERROR_U_P_NOT_MATCH)
		return
	}
	token, err := jwt.GenerateToken(username, password)
	if err != nil {
		appG.MakeEmptyResponse(http.StatusInternalServerError, e.ERROR_UNKNOWN)
		return
	}
	appG.SetCookie("ayapi_token", token, 7*24*3600, false, true)
	appG.MakeResponse(http.StatusOK, e.AUTH_OK, map[string]string{
		"token": token,
	})
}
