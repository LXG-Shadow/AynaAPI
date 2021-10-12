package anime

import (
	"AynaAPI/api/core"
	providerApi "AynaAPI/api/provider"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Resolve godoc
// @Summary resolve anime
// @Description 根据来源解析动漫
// @Tags Anime
// @Produce json
// @Param provider path string true "anime provider identifier"
// @Param url query string true "url"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "susudm?url=susudm-17891-acg-1, susudm?url=http://susudm.com/acg/17891/"
// @Router /api/v1/anime/resolve/{provider} [get]
func Resolve(context *gin.Context) {
	appG := app.AppGin{C: context}
	provider := context.Param("provider")
	if !providerApi.IsProviderAvailable(provider) {
		appG.MakeResponse(http.StatusBadRequest, e.BGM_PROVIDER_NOT_AVAILABLE, fmt.Sprintf("%s not avialable", provider))
		return
	}
	url, b := appG.C.GetQuery("url")
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require url")
		return
	}
	vModel, b := providerApi.InitWithUrl(provider, url)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_INVALID_PARAMETER, "url not match provider")
		return
	}
	if !api_service.ProviderInitialize(provider, vModel, useCache) {
		appG.MakeResponse(http.StatusOK, e.BGM_INITIALIZE_FAIL, "无法获取到对应id的信息")
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, struct {
		Info    *providerApi.ApiProvider `json:"info"`
		Playurl []core.ApiResource       `json:"playurl"`
	}{
		vModel,
		api_service.ProviderGetPlayUrls(vModel, useCache),
	})
}

// ResolveAll godoc
// @Summary resolve anime
// @Description 解析动漫
// @Tags Anime
// @Produce json
// @Param url query string true "url"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "susudm-17891-acg-1, http://susudm.com/acg/17891/"
// @Router /api/v1/anime/resolve [get]
func ResolveAll(context *gin.Context) {
	appG := app.AppGin{C: context}
	url, b := appG.C.GetQuery("url")
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require url")
		return
	}
	for provider, status := range providerApi.ProviderStatusMap {
		if status {
			vModel, b := providerApi.InitWithUrl(provider, url)
			if !b {
				continue
			}
			if !api_service.ProviderInitialize(provider, vModel, useCache) {
				continue
			}
			appG.MakeResponse(http.StatusOK, e.API_OK, struct {
				Info    *providerApi.ApiProvider `json:"info"`
				Playurl []core.ApiResource       `json:"playurl"`
			}{
				vModel,
				api_service.ProviderGetPlayUrls(vModel, useCache),
			})
			return
		}
	}
	appG.MakeResponse(http.StatusOK, e.BGM_NO_RESULT, "无法找到匹配的url")
}
