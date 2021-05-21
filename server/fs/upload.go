package fs

import (
	"AynaAPI/config"
	"fmt"
	"regexp"
)

func GetUploadPath() string {
	return config.ServerConfig.GetFilePath("upload")
}

func GetUploadUrl() string {
	return "/upload"
}

func GetUploadFileUrl(filepath string) string {
	return fmt.Sprintf("%s/%s", GetUploadUrl(), filepath)
}

func GetUploadFileFullUrl(filepath string) string {
	return fmt.Sprintf("%s/static%s", config.ServerConfig.RealUrl, GetUploadFileUrl(filepath))
}

func IsImageContentType(contentType string) bool {
	return regexp.MustCompile("^image/").FindString(contentType) != ""
}
