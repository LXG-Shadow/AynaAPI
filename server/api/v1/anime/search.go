package anime

import (
	e2 "AynaAPI/api/core/e"
	providerApi "AynaAPI/api/provider"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"fmt"
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
// @Router /api/v1/anime/search/{provider} [get]
func Search(context *gin.Context) {
	appG := app.AppGin{C: context}
	provider := context.Param("provider")
	if !providerApi.IsProviderAvailable(provider) {
		appG.MakeResponse(http.StatusBadRequest, e.BGM_PROVIDER_NOT_AVAILABLE, fmt.Sprintf("%s not avialable", provider))
		return
	}
	keyword, b := appG.C.GetQuery("keyword")
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require keyword")
		return
	}
	page := appG.GetIntQueryWithDefault("page", 1)
	result := api_service.ProviderSearch(provider, keyword, page, useCache)
	if result.Status != e2.SUCCESS {
		appG.MakeResponse(http.StatusOK, e.BGM_SEARCH_FAIL, result)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, result.Data)
}

// SearchAll godoc
// @Summary search anime
// @Description 搜索动漫
// @Tags Anime
// @Produce json
// @Param keyword query string true "keyword"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "刀剑神域"
// @Router /api/v1/anime/search [get]
func SearchAll(context *gin.Context) {
	appG := app.AppGin{C: context}
	keyword, b := appG.C.GetQuery("keyword")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require keyword")
		return
	}
	page := appG.GetIntQueryWithDefault("page", 1)
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	result := map[string]map[string]interface{}{}
	for provider, status := range providerApi.ProviderStatusMap {
		if status {
			resp := api_service.ProviderSearch(provider, keyword, page, useCache)
			if resp.Status == e2.SUCCESS {
				result[provider] = resp.Data
			}
		}
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, result)
}
