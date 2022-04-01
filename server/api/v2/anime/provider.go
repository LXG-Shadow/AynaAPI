package anime

import (
	"AynaAPI/api/anime"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProviderList godoc
// @Summary get provider list
// @Description 获取来源列表
// @Tags Anime
// @Produce json
// @Success 200 {object} resp.AnimeProviderList
// @Router /api/v2/anime/providerlist [get]
func GetProviderList(context *gin.Context) {
	appG := app.AppGin{C: context}
	appG.MakeResponse(http.StatusOK, e.API_OK, anime.Providers.GetProviderList())
	return
}
