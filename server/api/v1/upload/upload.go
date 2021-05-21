package upload

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/fs"
	"AynaAPI/utils/vfile"
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
	contentType, err := vfile.GetFileHeaderContentType(fh)
	if !fs.IsImageContentType(contentType) {
		appG.MakeEmptyResponse(http.StatusBadRequest, e.API_ERROR_UPLOAD_IVALID_IMAGE)
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
