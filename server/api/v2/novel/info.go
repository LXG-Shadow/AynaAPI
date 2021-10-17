package novel

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetInfo godoc
// @Summary get novel info
// @Description 获取小说简介
// @Tags Novel
// @Produce json
// @Param mid query string true "novel provider meta id"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse ""
// @Router /api/v2/novel/info [get]
func GetInfo(context *gin.Context) {
	appG := app.AppGin{C: context}
	metadata, b := appG.C.GetQuery("mid")
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require mid")
		return
	}
	result, errcode := api_service.NovelGet(metadata, useCache)
	if errcode != 0 {
		appG.MakeEmptyResponse(http.StatusOK, errcode)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, result)
	return
}
