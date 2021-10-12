package novel

import (
	e2 "AynaAPI/api/core/e"
	novelApi "AynaAPI/api/novel"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetContent godoc
// @Summary get novel content
// @Description 获取小说章节内容
// @Tags Novel
// @Produce json
// @Param url query string true "content url"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "https://www.linovelib.com/novel/2342/133318.html"
// @Router /api/v1/novel/content [get]
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
	if resp.Status != e2.SUCCESS {
		appG.MakeResponse(http.StatusOK, e.NOVEL_GET_DATA_FAIL, resp)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, resp.Data)
	return
}
