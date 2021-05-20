package vfile

import (
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
