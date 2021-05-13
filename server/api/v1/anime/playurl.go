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
		appG.MakeResponse(http.StatusBadRequest, e.API_REQUIRE_PARAMETER, "require uid")
		return
	}
	vModel := *providerApi.InitWithUid(provider, uid)
	if vModel == nil {
		appG.MakeResponse(http.StatusBadRequest, e.API_INVALID_PARAMETER, "uid not valid")
		return
	}
	if !vModel.Initialize() {
		appG.MakeResponse(http.StatusOK, e.BGM_INITIALIZE_FAIL, "无法获取到对应id的信息")
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, vModel.GetPlayUrls())
}
