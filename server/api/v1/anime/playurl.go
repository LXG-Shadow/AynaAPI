package anime

import (
	providerApi "AynaAPI/api/provider"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPlayUrl(context *gin.Context) {
	appG := app.AppGin{C: context}
	provider := context.Param("provider")
	if !providerApi.IsProviderAvailable(provider) {
		appG.MakeResponse(http.StatusBadRequest, e.BGM_PROVIDER_NOT_AVAILABLE, fmt.Sprintf("%s not avialable", provider))
		return
	}
	uid, b := appG.C.GetQuery("uid")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require uid")
		return
	}
	vModel, b := providerApi.InitWithUid(provider, uid)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_INVALID_PARAMETER, "uid not valid")
		return
	}
	if !(*vModel).Initialize() {
		appG.MakeResponse(http.StatusOK, e.BGM_INITIALIZE_FAIL, "无法获取到对应id的信息")
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, (*vModel).GetPlayUrls())
}

func GetPlayUrlAll(context *gin.Context) {
	appG := app.AppGin{C: context}
	uid, b := appG.C.GetQuery("uid")
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
			if !(*vModel).Initialize() {
				continue
			}
			appG.MakeResponse(http.StatusOK, e.API_OK, (*vModel).GetPlayUrls())
			return
		}
	}
	appG.MakeResponse(http.StatusOK, e.BGM_NO_RESULT, "无法找到匹配的uid")
}
