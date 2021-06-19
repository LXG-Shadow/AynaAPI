package novel

import (
	novelApi "AynaAPI/api/novel"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProviderList(context *gin.Context) {
	appG := app.AppGin{C: context}
	plist := make([]string, 0)
	for id, _ := range novelApi.ProviderMap {
		plist = append(plist, id)
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, plist)
	return
}

func GetProviderRules(context *gin.Context) {
	appG := app.AppGin{C: context}
	rlist := map[string]interface{}{}
	for id, info := range novelApi.ProviderMap {
		rlist[id] = info
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, rlist)
}
