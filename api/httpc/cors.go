package httpc

import (
	"AynaAPI/utils/vhttp"
)

func GetCORS(uri string, header map[string]string) string {
	host := vhttp.GetUrlHost(uri)
	return Get(uri, map[string]string{
		"origin":     host,
		"referer":    host,
		"user-agent": GetRandomUserAgent(),
	}).String()
}
