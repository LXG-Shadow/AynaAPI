package music

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetInfo godoc
// @Summary get music info
// @Description 根据metadata获取音乐信息
// @Tags Music
// @Produce json
// @Param mid query string true "music provider meta id"
// @Param cache query boolean false "use cache"
// @Success 200 {object} resp.MusicInfo
// @Router /api/v2/music/info [get]
func GetInfo(context *gin.Context) {
	appG := app.AppGin{C: context}
	metadata, b := appG.C.GetQuery("mid")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require mid")
		return
	}
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	result, errcode := api_service.MusicGet(metadata, useCache)
	if errcode != 0 {
		appG.MakeEmptyResponse(http.StatusOK, errcode)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, result)
}
