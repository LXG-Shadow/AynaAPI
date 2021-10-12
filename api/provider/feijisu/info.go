package susudm

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
	playerApi = Host + "/%s/%s/%s.html"
	infoApi   = Host + "/%s/%s/"
	dataApi   = "http://d.gqyy8.com:8077/ne2/s%s.js"
)

func GetPlayerApi(id string, category string, epId string) string {
	return fmt.Sprintf(playerApi, category, id, epId)
}

func GetInfoApi(id string, category string) string {
	return fmt.Sprintf(infoApi, category, id)
}

func GetDataApi(id string) string {
	return fmt.Sprintf(dataApi, id)
}

func GetInfo(id string, category string, epId string) core.ApiResponse {
	result := httpc.Get(GetPlayerApi(id, category, epId), nil).String()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		return core.CreateEmptyApiResponseByStatus(e2.INTERNAL_ERROR)
	}
	title := doc.Find("div[class=tvinfo] > h3").Text()
	if title == "" {
		return core.CreateEmptyApiResponseByStatus(e2.EXTERNAL_API_ERROR)
	}
	result = httpc.Get(GetInfoApi(id, category), nil).String()
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		return core.CreateEmptyApiResponseByStatus(e2.INTERNAL_ERROR)
	}
	eps := make([]string, 0)
	hrefExp := regexp.MustCompile("[0-9]+\\.html")
	doc.Find("ul[class=details-con2-list]  > li > a").Each(func(i int, selection *goquery.Selection) {
		if href, b := selection.Attr("href"); b {
			if u := hrefExp.FindString(href); u != "" {
				eid, _ := utils.SliceString(u, 0, -5)
				eps = append(eps, GenerateUniqueId(id, category, eid))
			}
		}
	})
	pic, _ := doc.Find("div[class=pic] > img").Attr("data-original")
	return core.CreateApiResponseByStatus(e2.SUCCESS, map[string]interface{}{
		"title":    title,
		"pic":      pic,
		"episodes": eps,
	})
}

func GetPlayData(id string, epId string) core.ApiResponse {
	result := httpc.Get(GetDataApi(id), nil).String()
	urlsExp := regexp.MustCompile(fmt.Sprintf("playarr(_[0-9])?\\[%s\\]=\"[^\"]*\";", epId))
	urls := make([]string, 0)
	for _, url := range urlsExp.FindAllString(result, -1) {
		if i := strings.Index(url, "\""); i > 0 {
			if u, b := utils.SliceString(url, i+1, -2); b {
				ul := strings.Split(u, ",")
				urls = append(urls, ul[0])
			}
		}
	}
	return core.CreateApiResponseByStatus(e2.SUCCESS, map[string]interface{}{
		"urls": urls,
	})
}
