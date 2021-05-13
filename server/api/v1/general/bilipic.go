package general

import (
	apiE "AynaAPI/api/model/e"
	"AynaAPI/api/uf/imghost"
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadBiliPic(context *gin.Context) {
	appG := app.AppGin{C: context}
	b64val, err := appG.C.GetPostForm("cover")
	if !err {
		appG.MakeEmptyResponse(http.StatusBadRequest, e.API_REQUIRE_PARAMETER)
		return
	}
	resp := imghost.UploadBilibili(b64val)
	if resp.Status != apiE.SUCCESS {
		appG.MakeResponse(http.StatusOK, e.API_ERROR_UNKNOWN, resp)
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, resp)
}
