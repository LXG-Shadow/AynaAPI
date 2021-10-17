package anime

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetInfo godoc
// @Summary get anime info
// @Description 根据metadata获取动漫信息
// @Tags Anime
// @Produce json
// @Param mid query string true "anime provider meta id"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "xxxxxxx"
// @Router /api/v2/anime/info [get]
func GetInfo(context *gin.Context) {
	appG := app.AppGin{C: context}
	metadata, b := appG.C.GetQuery("mid")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require mid")
		return
	}
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	result, errcode := api_service.AnimeGet(metadata, useCache)
	if errcode != 0 {
		appG.MakeEmptyResponse(http.StatusOK, errcode)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, result)
}
