package anime

import (
	"AynaAPI/api/anime"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Search godoc
// @Summary search anime
// @Description 根据来源搜索动漫
// @Tags Anime
// @Produce json
// @Param provider path string true "anime provider identifier"
// @Param keyword query string true "keyword"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "susudm?keyword=刀剑神域"
// @Router /api/v2/anime/search/{provider} [get]
func Search(context *gin.Context) {
	appG := app.AppGin{C: context}
	providerName := context.Param("provider")
	keyword, b := appG.C.GetQuery("keyword")
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require keyword")
		return
	}
	result, errcode := api_service.AnimeSearch(providerName, keyword, useCache)
	if errcode != 0 {
		appG.MakeEmptyResponse(http.StatusOK, errcode)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, map[string][]anime.AnimeMeta{
		providerName: result.Result,
	})
}

// SearchAll godoc
// @Summary search anime
// @Description 搜索动漫
// @Tags Anime
// @Produce json
// @Param keyword query string true "keyword"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "刀剑神域"
// @Router /api/v2/anime/search [get]
func SearchAll(context *gin.Context) {
	appG := app.AppGin{C: context}
	keyword, b := appG.C.GetQuery("keyword")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require keyword")
		return
	}
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	result := map[string][]anime.AnimeMeta{}
	for _, providerName := range anime.Providers.GetProviderList() {
		if r, err := api_service.AnimeSearch(providerName, keyword, useCache); err == 0 {
			result[providerName] = r.Result
		}
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, result)
}
