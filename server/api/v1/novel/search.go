package novel

import (
	apiE "AynaAPI/api/model/e"
	novelApi "AynaAPI/api/novel"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	if resp.Status != apiE.SUCCESS {
		appG.MakeResponse(http.StatusOK, e.NOVEL_GET_DATA_FAIL, resp)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, resp.Data)
}

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
		fmt.Println(provider.Identifier, resp)
		if resp.Status == apiE.SUCCESS {
			result[provider.Identifier] = resp.Data
		}
	}
	fmt.Println(result)
	appG.MakeResponse(http.StatusOK, e.API_OK, result)
	return
}
