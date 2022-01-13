package general

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/utils/vhttp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// todo: rewrite use HandleContext

func BypassCorsFixFile(context *gin.Context) {
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
	path := "/"
	if uri.Path != "" {
		paths := strings.Split(uri.Path, "/")
		path += paths[len(paths)-1]
	}
	para := "?url=" + vhttp.QueryEscapeWithEncoding(value, "utf-8")
	if value, b := context.GetQuery("referer"); b {
		para += "&referer=" + vhttp.QueryEscapeWithEncoding(value, "utf-8")
	}
	if value, b := context.GetQuery("ua"); b {
		para += "&ua=" + vhttp.QueryEscapeWithEncoding(value, "utf-8")
	}
	context.Redirect(http.StatusTemporaryRedirect, context.Request.URL.Host+context.Request.URL.Path+path+para)
}

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

		if value, b := context.GetQuery("referer"); b {
			req.Header.Set("referer", value)
		}
		if value, b := context.GetQuery("ua"); b {
			fmt.Println(value)
			req.Header.Set("user-agent", value)
		}

	}
	proxy.ServeHTTP(context.Writer, context.Request)
}
