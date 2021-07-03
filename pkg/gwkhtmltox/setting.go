package gwkhtmltox

/*
#cgo CFLAGS: -I"C:/Program Files/wkhtmltopdf/include"
#cgo LDFLAGS: -L"C:/Program Files/wkhtmltopdf/bin" -lwkhtmltox
#include <stdlib.h>
#include <wkhtmltox/image.h>
*/
import "C"
import (
	"reflect"
	"strconv"
)

type LoadingSetting struct {
	Username             string
	Password             string
	JsDelay              uint64
	ZoomFactor           float64
	BlockLocalFileAccess bool
	StopSlowScript       bool
	DebugJavascript      bool
	LoadErrorHandling    string
	Proxy                string
	PrintMediaType       bool
}

type WebSetting struct {
	Background                 bool
	LoadImages                 bool
	EnableJavascript           bool
	EnableIntelligentShrinking bool
	MinimumFontSize            uint64
	DefaultEncoding            string
	UserStyleSheet             string
	EnablePlugins              bool
}

type ImageSetting struct {
	CropLeft    uint64
	CropTop     uint64
	CropWidth   uint64
	CropHeight  uint64
	CookieJar   string
	Load        LoadingSetting
	Web         WebSetting
	Transparent bool
	Fmt         string
	ScreenWidth uint64
	SmartWidth  bool
	Quality     uint64
}

func getString(v interface{}) (string, bool) {
	switch reflect.TypeOf(v).Name() {
	case "uint64":
		val := v.(uint64)
		return strconv.FormatUint(val, 10), val > 0
	case "string":
		val := v.(string)
		return val, val != ""
	case "bool":
		val := v.(bool)
		return strconv.FormatBool(val), true
	default:
		return "", false
	}
}
