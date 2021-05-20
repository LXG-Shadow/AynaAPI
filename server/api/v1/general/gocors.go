package general

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/utils/vhttp"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func BypassCors(context *gin.Context) {
	appG := app.AppGin{C: context}
	value, b := context.GetQuery("url")
	if !b {
		appG.MakeResponse(http.StatusBadRequest,
			e.API_ERROR_REQUIRE_PARAMETER, "url required")
		return
	}
	if !vhttp.IsUrl(value) {
		appG.MakeResponse(http.StatusBadRequest,
			e.API_ERROR_INVALID_PARAMETER, "not proper url")
		return
	}
	uri, _ := url.Parse(value)
	proxy := httputil.NewSingleHostReverseProxy(uri)
	proxy.Director = func(req *http.Request) {
		req.URL = uri
		req.Host = uri.Host
		req.Header.Set("referer", uri.Host)
		req.Header.Set("origin", uri.Host)
		req.RemoteAddr = ""
		req.RequestURI = ""

		value, b := context.GetQuery("referer")
		if b {
			req.Header.Set("referer", value)
		}
	}
	proxy.ServeHTTP(context.Writer, context.Request)
}
