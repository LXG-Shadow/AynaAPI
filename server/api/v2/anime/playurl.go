package anime

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPlayUrl godoc
// @Summary get anime playurl
// @Description 获取动漫播放地址
// @Tags Anime
// @Produce json
// @Param mid query string true "anime provider meta id"
// @Param pid query int false "playlist id"
// @Param eid query int false "episode id"
// @Param cache query boolean false "use cache"
// @Success 200 {object} resp.AnimePlayUrl
// @Router /api/v2/anime/playurl [get]
func GetPlayUrl(context *gin.Context) {
	appG := app.AppGin{C: context}
	metadata, b := appG.C.GetQuery("mid")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require mid")
		return
	}
	pid := appG.GetIntQueryWithDefault("pid", 0)
	eid := appG.GetIntQueryWithDefault("eid", 0)
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	resp, errcode := api_service.AnimeGetVideo(metadata, pid, eid, useCache)
	if errcode != 0 {
		appG.MakeEmptyResponse(http.StatusOK, errcode)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, resp)
}
