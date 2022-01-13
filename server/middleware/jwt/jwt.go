package jwt

import (
	"AynaAPI/config"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/common"
	"AynaAPI/server/service/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthUser is jwt middleware
func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.AppGin{C: c}
		var token string
		if t, b := appG.C.GetQuery("token"); b {
			token = t
		} else {
			if cookie, err := appG.C.Cookie(config.ServerConfig.JwtTokenName); err == nil {
				token = cookie
			}
		}
		if token == "" {
			appG.MakeEmptyResponse(http.StatusUnauthorized, e.API_ERROR_REQUIRE_TOKEN)
			c.Abort()
			return
		}
		var claim *Claims
		if claim, _ := ParseToken(token); claim == nil {
			appG.MakeEmptyResponse(http.StatusUnauthorized, e.API_ERROR_INVALID_TOKEN)
			c.Abort()
			return
		}

		if user, ok := auth_service.New().AuthUser(auth_service.LoginParam{Username: claim.Username, Password: claim.Password}); !ok {
			appG.MakeEmptyResponse(http.StatusUnauthorized, e.API_ERROR_INVALID_TOKEN)
			c.Abort()
			return
		} else {
			appG.SetUser(user)
			c.Next()
		}
	}
}

func GetUserOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.AppGin{C: c}
		var token string
		if t, b := appG.C.GetQuery("token"); b {
			token = t
		} else {
			if cookie, err := appG.C.Cookie(config.ServerConfig.JwtTokenName); err == nil {
				token = cookie
			}
		}
		if token == "" {
			c.Next()
			return
		}
		var claim *Claims
		if claim, _ = ParseToken(token); claim == nil {
			c.Next()
			return
		}
		common.Logger.Info(claim)
		if user, ok := auth_service.New().AuthUser(auth_service.LoginParam{Username: claim.Username, Password: claim.Password}); !ok {
			c.Next()
			return
		} else {
			appG.SetUser(user)
			c.Next()
		}
	}
}
