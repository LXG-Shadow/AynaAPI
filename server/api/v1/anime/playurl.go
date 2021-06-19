package anime

import (
	providerApi "AynaAPI/api/provider"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPlayUrl godoc
// @Summary get anime playurl
// @Description 根据来源获取动漫播放地址
// @Tags Anime
// @Produce json
// @Param provider path string true "anime provider identifier"
// @Param uid query string true "uid"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "susudm?uid=susudm-17891-acg-1"
// @Router /api/v1/anime/playurl/{provider} [get]
func GetPlayUrl(context *gin.Context) {
	appG := app.AppGin{C: context}
	provider := context.Param("provider")
	if !providerApi.IsProviderAvailable(provider) {
		appG.MakeResponse(http.StatusBadRequest, e.BGM_PROVIDER_NOT_AVAILABLE, fmt.Sprintf("%s not avialable", provider))
		return
	}
	uid, b := appG.C.GetQuery("uid")
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require uid")
		return
	}
	vModel, b := providerApi.InitWithUid(provider, uid)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_INVALID_PARAMETER, "uid not valid")
		return
	}
	if !api_service.ProviderInitialize(provider, vModel, useCache) {
		appG.MakeResponse(http.StatusOK, e.BGM_INITIALIZE_FAIL, "无法获取到对应id的信息")
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, api_service.ProviderGetPlayUrls(vModel, useCache))
}

// GetPlayUrlAll godoc
// @Summary get anime playurl
// @Description 获取动漫播放地址
// @Tags Anime
// @Produce json
// @Param uid query string true "uid"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "susudm-17891-acg-1"
// @Router /api/v1/anime/playurl [get]
func GetPlayUrlAll(context *gin.Context) {
	appG := app.AppGin{C: context}
	uid, b := appG.C.GetQuery("uid")
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require uid")
		return
	}
	for provider, status := range providerApi.ProviderStatusMap {
		if status {
			vModel, b := providerApi.InitWithUid(provider, uid)
			if !b {
				continue
			}
			if !api_service.ProviderInitialize(provider, vModel, useCache) {
				continue
			}
			appG.MakeResponse(http.StatusOK, e.API_OK, api_service.ProviderGetPlayUrls(vModel, useCache))
			return
		}
	}
	appG.MakeResponse(http.StatusOK, e.BGM_NO_RESULT, "无法找到匹配的uid")
}
