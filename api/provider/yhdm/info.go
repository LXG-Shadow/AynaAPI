package yhdm

import (
	"AynaAPI/api/core"
	e2 "AynaAPI/api/core/e"
	"AynaAPI/api/httpc"
	"AynaAPI/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

const (
	playerUrl       = Host + "/v/%s-%s.html"
	resolveVidApi   = "http://tup.yhdm.so/?vid=%s"
	resolveQzoneApi = "http://tup.yhdm.so/qzone.php?url=%s"
)

func GetPlayerApi(id string, epId string) string {
	return fmt.Sprintf(playerUrl, id, epId)
}

func GetResolveApi(rawurl string) string {
	vfmt := strings.Split(rawurl, "$")
	if len(vfmt) < 2 {
		return fmt.Sprintf(resolveVidApi, rawurl)
	}
	switch vfmt[1] {
	case "mp4":
		return fmt.Sprintf(resolveVidApi, rawurl)
	case "qzz":
		return fmt.Sprintf(resolveQzoneApi, vfmt[0])
	default:
		return fmt.Sprintf(resolveVidApi, rawurl)
	}
}

func GetInfo(id string, epId string) core.ApiResponse {
	result := httpc.Get(GetPlayerApi(id, epId), nil).String()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		return core.CreateEmptyApiResponseByStatus(e2.INTERNAL_ERROR)
	}
	title := doc.Find("h1 > a").Text()

	if title == "" {
		return core.CreateEmptyApiResponseByStatus(e2.EXTERNAL_API_ERROR)
	}

	eps := make([]string, 0)
	hrefExp := regexp.MustCompile("[0-9]+-[^-]*\\.html")
	doc.Find("div[class=movurls] > ul > li > a").Each(func(i int, selection *goquery.Selection) {
		if hLink, b := selection.Attr("href"); b {
			if s := hrefExp.FindString(hLink); s != "" {
				idSL := strings.Split(strings.ReplaceAll(s, ".html", ""), "-")
				eps = append(eps, GenerateUniqueId(idSL[0], idSL[1]))
			}
		}
	})
	playdataUrl, b := doc.Find("div[id=playbox]").Attr("data-vid")

	if !b {
		return core.CreateEmptyApiResponseByStatus(e2.EXTERNAL_API_ERROR)
	}
	pic, _ := utils.SliceString(regexp.MustCompile("var bdPic = \"[^;]*\";").FindString(result), 13, -2)

	return core.CreateApiResponseByStatus(e2.SUCCESS, map[string]interface{}{
		"title":    title,
		"pic":      pic,
		"playUrl":  playdataUrl,
		"episodes": eps,
	})
}

func ResolveVideoUrl(url string) core.ApiResponse {
	result := httpc.Get(GetResolveApi(url), nil).String()
	realUrl, b := utils.SliceString(regexp.MustCompile("url: \"(.*)\",").FindString(result), 6, -2)
	if !b {
		return core.CreateEmptyApiResponseByStatus(e2.EXTERNAL_API_ERROR)
	}
	return core.CreateApiResponseByStatus(e2.SUCCESS, map[string]interface{}{
		"url":     url,
		"realUrl": realUrl,
	})
}
