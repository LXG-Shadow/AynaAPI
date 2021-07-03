package anime

import (
	animeApi "AynaAPI/api/provider"
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
// @Router /api/v1/anime/providerlist [get]
func GetProviderList(context *gin.Context) {
	appG := app.AppGin{C: context}
	plist := make([]map[string]string, 0)
	for id, val := range animeApi.ProviderStatusMap {
		if val {
			plist = append(plist, map[string]string{"id": id, "name": id})
		}
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, plist)
	return
}
