package vfile

import (
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

func GetFileExt(filename string) string {
	return path.Ext(filename)
}

func GetFileName(filename string) string {
	ext := path.Ext(filename)
	return strings.TrimSuffix(filename, ext)
}

func GetFileHeaderContentType(file *multipart.FileHeader) (string, error) {
	out, err := file.Open()
	if err != nil {
		return "", err
	}

	buffer := make([]byte, 512)

	_, err = out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
