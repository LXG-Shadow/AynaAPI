package general

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/fs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(c *gin.Context) {
	appG := app.AppGin{C: c}
	fh, err := c.FormFile("image")
	if err != nil {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_UPLOAD_IMAGE_NOT_FOUND, "require image in post form")
		return
	}
	filename, err := appG.SaveUploadedFileWithMD5(fh)
	if err != nil {
		appG.MakeEmptyResponse(http.StatusInternalServerError, e.API_ERROR_UPLOAD_SAVE_IMAGE_FAIL)
		return
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, map[string]string{
		"filename": filename,
		"url":      fs.GetUploadFileUrl(filename),
		"real_url": fs.GetUploadFileFullUrl(filename),
	})
}
