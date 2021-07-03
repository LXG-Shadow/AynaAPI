package novel

import (
	novelApi "AynaAPI/api/novel"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProviderList godoc
// @Summary get provider list
// @Description 获取来源列表
// @Tags Novel
// @Produce json
// @Success 200 {object} app.AppJsonResponse
// @Router /api/v1/novel/providerlist [get]
func GetProviderList(context *gin.Context) {
	appG := app.AppGin{C: context}
	plist := make([]map[string]string, 0)
	for id, val := range novelApi.ProviderMap {
		plist = append(plist, map[string]string{"id": id, "name": val.Name})
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, plist)
	return
}

// GetProviderRules godoc
// @Summary get provider rules
// @Description 获取来源规则
// @Tags Novel
// @Produce json
// @Success 200 {object} app.AppJsonResponse
// @Router /api/v1/novel/rulelist [get]
func GetProviderRules(context *gin.Context) {
	appG := app.AppGin{C: context}
	rlist := map[string]interface{}{}
	for id, info := range novelApi.ProviderMap {
		rlist[id] = info
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, rlist)
}
