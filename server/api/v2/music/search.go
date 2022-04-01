package music

import (
	"AynaAPI/api/music"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Search godoc
// @Summary search music
// @Description 根据来源搜索音乐
// @Tags Music
// @Produce json
// @Param provider path string true "music provider identifier (e.g. bilibilimusic)"
// @Param keyword query string true "keyword (e.g. 霜雪千年)"
// @Param cache query boolean false "use cache"
// @Success 200 {object} resp.MusicSearchResult
// @Router /api/v2/music/search/{provider} [get]
func Search(context *gin.Context) {
	appG := app.AppGin{C: context}
	providerName := context.Param("provider")
	keyword, b := appG.C.GetQuery("keyword")
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require keyword")
		return
	}
	result, errcode := api_service.MusicSearch(providerName, keyword, useCache)
	if errcode != 0 {
		appG.MakeEmptyResponse(http.StatusOK, errcode)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, map[string][]music.MusicMeta{
		providerName: result.Result,
	})
}

// SearchAll godoc
// @Summary search music
// @Description 搜索音乐
// @Tags Music
// @Produce json
// @Param keyword query string true "keyword"
// @Param cache query boolean false "use cache"
// @Success 200 {object} resp.MusicSearchResult
// @Router /api/v2/music/search [get]
func SearchAll(context *gin.Context) {
	appG := app.AppGin{C: context}
	keyword, b := appG.C.GetQuery("keyword")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require keyword")
		return
	}
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	result := map[string][]music.MusicMeta{}
	for _, providerName := range music.Providers.GetProviderList() {
		if r, err := api_service.MusicSearch(providerName, keyword, useCache); err == 0 {
			result[providerName] = r.Result
		}
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, result)
}
