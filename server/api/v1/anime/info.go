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

// GetInfo godoc
// @Summary get anime info
// @Description 根据来源获取动漫信息
// @Tags Anime
// @Produce json
// @Param provider path string true "anime provider identifier"
// @Param uid query string true "uid"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "susudm?uid=susudm-17891-acg-1"
// @Router /api/v1/anime/info/{provider} [get]
func GetInfo(context *gin.Context) {
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
	appG.MakeResponse(http.StatusOK, e.API_OK, vModel)
}

// InfoAll godoc
// @Summary get anime info
// @Description 获取动漫信息
// @Tags Anime
// @Produce json
// @Param uid query string true "uid"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "susudm-17891-acg-1"
// @Router /api/v1/anime/info [get]
func InfoAll(context *gin.Context) {
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
			appG.MakeResponse(http.StatusOK, e.API_OK, vModel)
			return
		}
	}
	appG.MakeResponse(http.StatusOK, e.BGM_NO_RESULT, "无法找到匹配的uid")
}
