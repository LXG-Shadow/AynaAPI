package novel

import (
	e2 "AynaAPI/api/core/e"
	novelApi "AynaAPI/api/novel"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"fmt"
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
// @Router /api/v1/novel/search/{provider} [get]
func Search(context *gin.Context) {
	appG := app.AppGin{C: context}
	providerName := context.Param("provider")
	if !novelApi.IsProviderAvailable(providerName) {
		appG.MakeResponse(http.StatusBadRequest, e.NOVEL_PROVIDER_NOT_AVAILABLE, fmt.Sprintf("%s not avialable", providerName))
		return
	}
	keyword, b := appG.C.GetQuery("keyword")
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require keyword")
		return
	}
	provider := novelApi.GetNovelProvider(providerName)
	resp := api_service.NovelSearch(provider, keyword, useCache)
	if resp.Status != e2.SUCCESS {
		appG.MakeResponse(http.StatusOK, e.NOVEL_GET_DATA_FAIL, resp)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, resp.Data)
}

// SearchAll godoc
// @Summary search novel
// @Description 搜索小说
// @Tags Novel
// @Produce json
// @Param keyword query string true "keyword"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "诡秘之主"
// @Router /api/v1/novel/search [get]
func SearchAll(context *gin.Context) {
	appG := app.AppGin{C: context}
	keyword, b := appG.C.GetQuery("keyword")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require keyword")
		return
	}
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	result := map[string]map[string]interface{}{}
	for _, provider := range novelApi.ProviderMap {
		resp := api_service.NovelSearch(&provider, keyword, useCache)
		if resp.Status == e2.SUCCESS {
			result[provider.Identifier] = resp.Data
		}
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, result)
	return
}
