package upload

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/server/common"
	"AynaAPI/server/pkg/upload"
	"AynaAPI/server/service/web_service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

type fileInfoWeb struct {
	FileHash    string `json:"file_hash"`
	ContentType string `json:"content_type"`
	AccessUrl   string `json:"access_url"`
}

func UploadFile(c *gin.Context) {
	appG := app.AppGin{C: c}
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		appG.MakeErrorResponse(http.StatusOK, e.API_ERROR_INVALID_PARAMETER, err.Error())
		return
	}
	fileType := cast.ToInt(c.PostForm("type"))
	if fileHeader == nil || fileType <= 0 {
		appG.MakeErrorResponse(http.StatusOK, e.API_ERROR_INVALID_PARAMETER)
		return
	}

	svc := web_service.New()
	uploadFile, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		common.Logger.WithContext(c).Errorf("svc.UploadFile err: %v", err)
		appG.MakeErrorResponse(http.StatusOK, e.API_ERROR_UPLOAD_SAVE_IMAGE_FAIL, err.Error())
		return
	}

	appG.MakeResponse(http.StatusOK, e.API_OK, fileInfoWeb{
		FileHash:    uploadFile.FileHash,
		ContentType: uploadFile.ContentType,
		AccessUrl:   uploadFile.AccessUrl,
	})
}
