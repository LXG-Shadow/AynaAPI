package novel

import (
	"AynaAPI/api/novel"
	novelCore "AynaAPI/api/novel/core"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Search godoc
// @Summary search novel
// @Description 根据来源搜索小说
// @Tags Novel
// @Produce json
// @Param provider path string true "novel provider identifier"
// @Param keyword query string true "keyword"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "biqugeB?keyword=诡秘之主"
// @Router /api/v2/novel/search/{provider} [get]
func Search(context *gin.Context) {
	appG := app.AppGin{C: context}
	providerName := context.Param("provider")
	keyword, b := appG.C.GetQuery("keyword")
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require keyword")
		return
	}
	resp, errcode := api_service.NovelSearch(providerName, keyword, useCache)
	if errcode != 0 {
		appG.MakeEmptyResponse(http.StatusOK, errcode)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, resp)
}

// SearchAll godoc
// @Summary search novel
// @Description 搜索小说
// @Tags Novel
// @Produce json
// @Param keyword query string true "keyword"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "诡秘之主"
// @Router /api/v2/novel/search [get]
func SearchAll(context *gin.Context) {
	appG := app.AppGin{C: context}
	keyword, b := appG.C.GetQuery("keyword")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require keyword")
		return
	}
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	result := map[string]novelCore.NovelSearchResult{}
	for _, providerName := range novel.GetNovelProviderList() {
		if r, err := api_service.NovelSearch(providerName, keyword, useCache); err == 0 {
			result[providerName] = r
		}
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, result)
	return
}
