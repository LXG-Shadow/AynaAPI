package anime

import (
	apiE "AynaAPI/api/model/e"
	providerApi "AynaAPI/api/provider"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search(context *gin.Context) {
	appG := app.AppGin{C: context}
	provider := context.Param("provider")
	if !providerApi.IsProviderAvailable(provider) {
		appG.MakeResponse(http.StatusBadRequest, e.BGM_PROVIDER_NOT_AVAILABLE, fmt.Sprintf("%s not avialable", provider))
		return
	}
	keyword, b := appG.C.GetQuery("keyword")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_REQUIRE_PARAMETER, "require keyword")
		return
	}
	page := appG.GetIntQueryWithDefault("page", 1)
	result := providerApi.Search(provider, keyword, page)
	if result.Status != apiE.SUCCESS {
		appG.MakeResponse(http.StatusOK, e.BGM_SEARCH_FAIL, result)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, result.Data)
}

func SearchAll(context *gin.Context) {
	appG := app.AppGin{C: context}
	keyword, b := appG.C.GetQuery("keyword")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_REQUIRE_PARAMETER, "require keyword")
		return
	}
	page := appG.GetIntQueryWithDefault("page", 1)
	result := map[string]map[string]interface{}{}
	for provider, status := range providerApi.ProviderStatusMap {
		if status {
			resp := providerApi.Search(provider, keyword, page)
			if resp.Status == apiE.SUCCESS {
				result[provider] = resp.Data
			}
		}
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, result)
}
