package anime

import (
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
	appG.MakeResponse(http.StatusOK, e.API_OK, result)
}
