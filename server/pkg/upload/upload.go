package upload

import (
	"AynaAPI/config"
	"AynaAPI/utils/vfile"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

func GetFullSavePath(filename string) string {
	return config.ServerConfig.UploadSavePath + "/" + filename
}

func GetFullAccessUrl(filename string) string {
	return config.ServerConfig.UploadServerUrl + "/" + filename
}

func GetMD5FileInfo(fileHeader *multipart.FileHeader) (string, string, bool) {
	ext := vfile.GetFileExt(fileHeader.Filename)
	md5Str, err := vfile.CalcFileHeaderMD5(fileHeader)
	if err != nil {
		return "", "", false
	}
	return md5Str, ext, true
}

func GetFileContentType(fileHeader *multipart.FileHeader) (string, bool) {
	contentType, err := vfile.GetFileHeaderContentType(fileHeader)
	if err != nil {
		return "", false
	}
	return contentType, true
}

func GetSavePath() string {
	return config.ServerConfig.UploadSavePath
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CheckContainExt(t FileType, ext string) bool {
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range config.ServerConfig.UploadAllowImageExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= config.ServerConfig.UploadMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
