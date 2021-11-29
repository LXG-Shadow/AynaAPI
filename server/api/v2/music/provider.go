package music

import (
	"AynaAPI/api/music"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProviderList godoc
// @Summary get provider list
// @Description 获取来源列表
// @Tags Music
// @Produce json
// @Success 200 {object} app.AppJsonResponse
// @Router /api/v2/music/providerlist [get]
func GetProviderList(context *gin.Context) {
	appG := app.AppGin{C: context}
	appG.MakeResponse(http.StatusOK, e.API_OK, music.Providers.GetProviderList())
	return
}
