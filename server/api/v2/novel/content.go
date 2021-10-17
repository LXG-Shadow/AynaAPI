package novel

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/api_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetContent godoc
// @Summary get novel content
// @Description 获取小说章节内容
// @Tags Novel
// @Produce json
// @Param mid query string true "novel provider meta id"
// @Param vid query int false "volume id"
// @Param cid query int false "chapter id"
// @Param cache query boolean false "use cache"
// @Success 200 {object} app.AppJsonResponse "https://www.linovelib.com/novel/2342/133318.html"
// @Router /api/v2/novel/content [get]
func GetContent(context *gin.Context) {
	appG := app.AppGin{C: context}
	metadata, b := appG.C.GetQuery("mid")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "require url")
		return
	}
	vid := appG.GetIntQueryWithDefault("vid", 0)
	cid := appG.GetIntQueryWithDefault("cid", 0)
	useCache := appG.GetBoolQueryWithDefault("cache", true)
	resp, errcode := api_service.NovelContent(metadata, vid, cid, useCache)
	if errcode != 0 {
		appG.MakeEmptyResponse(http.StatusOK, errcode)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, resp)
	return
}
