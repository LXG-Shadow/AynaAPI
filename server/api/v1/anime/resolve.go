package anime

import (
	"AynaAPI/api/model"
	providerApi "AynaAPI/api/provider"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Resolve(context *gin.Context) {
	appG := app.AppGin{C: context}
	provider := context.Param("provider")
	if !providerApi.IsProviderAvailable(provider) {
		appG.MakeResponse(http.StatusBadRequest, e.BGM_PROVIDER_NOT_AVAILABLE, fmt.Sprintf("%s not avialable", provider))
		return
	}
	url, b := appG.C.GetQuery("url")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_REQUIRE_PARAMETER, "require url")
		return
	}
	vModel, b := providerApi.InitWithUrl(provider, url)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_INVALID_PARAMETER, "url not match provider")
		return
	}
	if !(*vModel).Initialize() {
		appG.MakeResponse(http.StatusOK, e.BGM_INITIALIZE_FAIL, "无法获取到对应id的信息")
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, struct {
		Info    *providerApi.ApiProvider `json:"info"`
		Playurl []model.ApiResource      `json:"playurl"`
	}{
		vModel,
		(*vModel).GetPlayUrls(),
	})
}
