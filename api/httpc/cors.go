package httpc

import (
	"AynaAPI/utils/vhttp"
)

func GetCORSString(uri string, header map[string]string) string {
	host := vhttp.GetUrlHost(uri)
	return GetBodyString(uri, map[string]string{
		"origin":     host,
		"referer":    host,
		"user-agent": GetRandomUserAgent(),
	})
}
