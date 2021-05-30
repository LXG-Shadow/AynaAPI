package susudm

import (
	"AynaAPI/api/httpc"
	"AynaAPI/api/model"
	"AynaAPI/api/model/e"
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

func GetInfo(id string, category string, epId string) model.ApiResponse {
	result := httpc.Get(GetPlayerApi(id, category, epId), nil).String()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	title := doc.Find("div[class=wrap] > div[class=h2] > a").Last().Text()
	if title == "" {
		return model.CreateEmptyApiResponseByStatus(e.EXTERNAL_API_ERROR)
	}
	result = httpc.Get(GetInfoApi(id, category), nil).String()
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	eps := make([]string, 0)
	hrefExp := regexp.MustCompile("[a-zA-Z]+/[0-9]+/[0-9]+")
	doc.Find("ul[class=urlli] > div > ul > li > a").Each(func(i int, selection *goquery.Selection) {
		if href, b := selection.Attr("href"); b {
			if u := hrefExp.FindString(href); u != "" {
				idlist := strings.Split(u, "/")
				eps = append(eps, GenerateUniqueId(idlist[1], idlist[0], idlist[2]))
			}
		}
	})
	pic, _ := doc.Find("div[class=pic] > img").Attr("data-original")
	return model.CreateApiResponseByStatus(e.SUCCESS, map[string]interface{}{
		"title":    title,
		"pic":      pic,
		"episodes": eps,
	})
}

func GetPlayData(id string, epId string) model.ApiResponse {
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
	return model.CreateApiResponseByStatus(e.SUCCESS, map[string]interface{}{
		"urls": urls,
	})
}
