package gwkhtmltox

/*
#cgo CFLAGS: -I"C:/Program Files/wkhtmltopdf/include"
#cgo LDFLAGS: -L"C:/Program Files/wkhtmltopdf/bin" -lwkhtmltox
#include <stdlib.h>
#include <wkhtmltox/image.h>
*/
import "C"
import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
	"unsafe"
)

// ImageInit initialize image library
func ImageInit() error {
	if C.wkhtmltoimage_init(0) != 1 {
		return errors.New("could not initialize library")
	}
	return nil
}

// ImageVersion Version returns the version of the library.
func ImageVersion() string {
	return C.GoString(C.wkhtmltoimage_version())
}

// ImageDestroy Destroy releases all the resources used by the library.
func ImageDestroy() {
	C.wkhtmltoimage_deinit()
}

func setImageSetting(settings *C.wkhtmltoimage_global_settings, name string, value interface{}) error {
	if name = strings.TrimSpace(name); name == "" {
		return errors.New("converter option name cannot be empty")
	}
	val, ok := getString(value)
	if !ok {
		return nil
	}
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	v := C.CString(val)
	defer C.free(unsafe.Pointer(v))

	if errCode := C.wkhtmltoimage_set_global_setting(settings, n, v); errCode != 1 {
		return fmt.Errorf("could not set converter option `%s` to `%s`: code %d", name, val, errCode)
	}
	return nil
}

type ImageConvertor struct {
	Setting *ImageSetting
}

func (self *ImageConvertor) parseImageSetting(settings *C.wkhtmltoimage_global_settings) error {
	s := map[string]interface{}{
		// image setting
		//"crop.left":      self.Setting.CropLeft,
		//"crop.top":       self.Setting.CropTop,
		//"crop.width":     self.Setting.CropWidth,
		//"crop.height":    self.Setting.CropHeight,
		//"load.cookieJar": self.Setting.CookieJar,
		"transparent": self.Setting.Transparent,
		"fmt":         self.Setting.Fmt,
		"screenWidth": self.Setting.ScreenWidth,
		//"smartWidth":     self.Setting.SmartWidth,
		"quality": self.Setting.Quality,
		// load setting
		//"load.username":             self.Setting.Load.Username,
		//"load.password":             self.Setting.Load.Password,
		//"load.jsdelay":              self.Setting.Load.JsDelay,
		//"load.zoomFactor":           self.Setting.Load.ZoomFactor,
		//"load.blockLocalFileAccess": self.Setting.Load.BlockLocalFileAccess,
		//"load.stopSlowScript":       self.Setting.Load.StopSlowScript,
		//"load.debugJavascript":      self.Setting.Load.DebugJavascript,
		//"load.loadErrorHandling":    self.Setting.Load.LoadErrorHandling,
		//"load.proxy":                self.Setting.Load.Proxy,
		//"load.printMediaType":       self.Setting.Load.PrintMediaType,
		// web setting
		//"web.background":                 self.Setting.Web.Background,
		//"web.loadImages":                 self.Setting.Web.LoadImages,
		//"web.enableJavascript":           self.Setting.Web.EnableJavascript,
		//"web.enableIntelligentShrinking": self.Setting.Web.EnableIntelligentShrinking,
		//"web.minimumFontSize":            self.Setting.Web.MinimumFontSize,
		//"web.defaultEncoding":            self.Setting.Web.DefaultEncoding,
		//"web.userStyleSheet":             self.Setting.Web.UserStyleSheet,
		//"web.enablePlugins":              self.Setting.Web.EnablePlugins,
	}

	for key, val := range s {
		if err := setImageSetting(settings, key, val); err != nil {
			return err
		}
	}
	return nil
}

func (self *ImageConvertor) Convert(input string, w io.Writer) error {
	if w == nil {
		return errors.New("the provided writer cannot be nil")
	}

	// create new settings instance
	settings := C.wkhtmltoimage_create_global_settings()
	if settings == nil {
		return errors.New("could not create converter settings")
	}
	if err := self.parseImageSetting(settings); err != nil {
		return err
	}
	// create converter with global settings and input html
	converter := C.wkhtmltoimage_create_converter(settings, C.CString(input))
	if converter == nil {
		return errors.New("could not create converter")
	}

	// cleanup converter
	defer func() {
		C.wkhtmltoimage_destroy_converter(converter)
		converter = nil
	}()

	go func() {
		n := 0
		for true {
			fmt.Println(n)
			n++
			fmt.Println(C.GoString(C.wkhtmltoimage_progress_string(converter)))
			fmt.Println(C.GoString(C.wkhtmltoimage_progress_string(converter)))
			time.Sleep(time.Second * 2)
		}
	}()

	fmt.Println("start convert")
	// Convert objects.
	if C.wkhtmltoimage_convert(converter) != 1 {
		return errors.New("could not convert given html")
	}

	// Get conversion output buffer.
	var output *C.uchar
	size := C.wkhtmltoimage_get_output(converter, &output)
	if size == 0 {
		return errors.New("could not retrieve the converted file")
	}
	// Copy output to the provided writer.
	buf := bytes.NewBuffer(C.GoBytes(unsafe.Pointer(output), C.int(size)))
	if _, err := io.Copy(w, buf); err != nil {
		return err
	}
	return nil
}

func (self *ImageConvertor) ConvertToBytes(input string) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	err := self.Convert(input, buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
