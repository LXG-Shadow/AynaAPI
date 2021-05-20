package jwt

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.AppGin{C: c}
		var token string
		if t, b := appG.C.GetQuery("token"); b {
			token = t
		} else {
			if cookie, err := appG.C.Cookie("ayapi_token"); err == nil {
				token = cookie
			}
		}
		if token == "" {
			appG.MakeEmptyResponse(http.StatusUnauthorized, e.API_ERROR_REQUIRE_TOKEN)
			c.Abort()
			return
		} else {
			if claim, _ := ParseToken(token); claim != nil {
				if ok, _ := models.AuthUser(claim.Username, claim.Password); !ok {
					appG.MakeEmptyResponse(http.StatusUnauthorized, e.API_ERROR_INVALID_TOKEN)
					c.Abort()
					return
				}
			} else {
				appG.MakeEmptyResponse(http.StatusUnauthorized, e.API_ERROR_INVALID_TOKEN)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
