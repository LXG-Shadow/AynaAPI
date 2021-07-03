package app

import (
	"AynaAPI/server/app/e"
	"AynaAPI/server/fs"
	"AynaAPI/server/models"
	"AynaAPI/utils/vfile"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type AppGin struct {
	C *gin.Context
}

type AppJsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (g *AppGin) GetQueryWithDefault(key string, defaultValue string) string {
	val, b := g.C.GetQuery(key)
	if b {
		return val
	} else {
		return defaultValue
	}
}

func (g *AppGin) GetIntQueryWithDefault(key string, defaultValue int) int {
	if val, b := g.C.GetQuery(key); b {
		if ival, b1 := cast.ToIntE(val); b1 == nil {
			return ival
		}
	}
	return defaultValue
}

func (g *AppGin) GetBoolQueryWithDefault(key string, defaultValue bool) bool {
	if val, b := g.C.GetQuery(key); b {
		if bval, b1 := cast.ToBoolE(val); b1 == nil {
			return bval
		}
	}
	return defaultValue
}

func (g *AppGin) MakeResponse(httpCode int, statusCode int, data interface{}) {
	g.C.JSON(httpCode, AppJsonResponse{
		Code:    statusCode,
		Message: e.GetMessage(statusCode),
		Data:    data,
	})
	return
}

func (g *AppGin) MakeEmptyResponse(httpCode int, statusCode int) {
	g.C.JSON(httpCode, AppJsonResponse{
		Code:    statusCode,
		Message: e.GetMessage(statusCode),
		Data:    nil,
	})
	return
}

func (g *AppGin) SetCookie(name, value string, maxAge int, secure, httpOnly bool) {
	g.C.SetCookie(name, value, maxAge, "", "", secure, httpOnly)
}

func (g *AppGin) DeleteCookie(name string) {
	g.C.SetCookie(name, "", -1, "", "", true, true)
}

func (g *AppGin) SetUser(user *models.User) {
	g.C.Set("ayapi_user", user)
}

func (g *AppGin) GetUser() (user *models.User, exists bool) {
	if val, ok := g.C.Get("ayapi_user"); !ok {
		return nil, false
	} else {
		return val.(*models.User), true
	}
}

func (g *AppGin) SaveUploadedFileWithMD5(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	md5string, err := vfile.CalcFileHeaderMD5(file)
	if err != nil {
		return "", err
	}
	filename := md5string + vfile.GetFileExt(file.Filename)
	dst := filepath.Join(fs.GetUploadPath(), filename)
	out, err := os.Create(dst)
	if err != nil {
		return filename, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return filename, err
}
