package cookie

import (
	"AynaAPI/config"
	"net/http"
)

const (
	DEFAULT  = "default"
	BILIBILI = "bilibili"
)

var SiteCookieMap = map[string][]*http.Cookie{
	DEFAULT: []*http.Cookie{},
	BILIBILI: []*http.Cookie{
		{
			Name:  "SESSDATA",
			Value: config.APIConfig.Bilibili_SESSDATA,
		},
		{
			Name:  "bili_jct",
			Value: config.APIConfig.Bilibili_JCT,
		},
	},
}

func GetCookie(site string) []*http.Cookie {
	cookie, b := SiteCookieMap[site]
	if b {
		return cookie
	}
	return SiteCookieMap[DEFAULT]
}
