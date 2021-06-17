package novel

import (
	apiE "AynaAPI/api/model/e"
	novelApi "AynaAPI/api/novel"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetContent(context *gin.Context) {
	appG := app.AppGin{C: context}
	url, b := appG.C.GetQuery("url")
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require url")
		return
	}
	provider := novelApi.GetNovelProviderByUrl(url)
	if provider == nil {
		appG.MakeEmptyResponse(http.StatusOK, e.NOVEL_PROVIDER_NOT_AVAILABLE)
		return
	}
	if !provider.IsContentUrl(url) {
		appG.MakeEmptyResponse(http.StatusOK, e.NOVEL_URL_NOT_SUPPORT)
		return
	}
	resp := api_service.NovelContent(provider, url, useCache)
	if resp.Status != apiE.SUCCESS {
		appG.MakeResponse(http.StatusOK, e.NOVEL_GET_DATA_FAIL, resp)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, resp.Data)
	return
}
