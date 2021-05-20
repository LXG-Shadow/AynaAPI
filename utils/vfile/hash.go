package vfile

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
)

func CalcMD5(file []byte) string {
	md5Str := fmt.Sprintf("%x", md5.Sum(file))
	return md5Str
}

func CalcFileHeaderMD5(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func CalcFileMD5(filePath string) (string, error) {
	f, err := ioutil.ReadFile(filePath)
	if nil != err {
		return "", err
	}
	md5Str := fmt.Sprintf("%x", md5.Sum(f))
	return md5Str, nil
}
