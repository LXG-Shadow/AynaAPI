package music

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUrl godoc
// @Summary get music playurl
// @Description 获取音乐播放地址
// @Tags Music
// @Produce json
// @Param mid query string true "music provider meta id"
// @Param ua query string false "specify user agent"
// @Param cache query boolean false "use cache"
// @Success 200 {object} resp.MusicPlayUrl
// @Router /api/v2/music/url [get]
func GetUrl(context *gin.Context) {
	appG := app.AppGin{C: context}
	metadata, b := appG.C.GetQuery("mid")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require mid")
		return
	}
	ua := appG.GetQueryWithDefault("ua", appG.C.Request.UserAgent())
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	resp, errcode := api_service.MusicGetAudio(metadata, ua, useCache)
	if errcode != 0 {
		appG.MakeEmptyResponse(http.StatusOK, errcode)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, resp)
}
