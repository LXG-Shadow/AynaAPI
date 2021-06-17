package httpc

import (
	"AynaAPI/utils/vhttp"
	"github.com/levigross/grequests"
)

func GetCORS(uri string) *grequests.Response {
	host := vhttp.GetUrlHost(uri)
	return Get(uri, map[string]string{
		"origin":     host,
		"referer":    host,
		"user-agent": GetRandomUserAgent(),
	})
}
