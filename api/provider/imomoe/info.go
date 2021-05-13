package imomoe

import (
	"AynaAPI/api/httpc"
	"AynaAPI/api/model"
	"AynaAPI/api/model/e"
	"AynaAPI/utils"
	"AynaAPI/utils/vhttp"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"regexp"
	"strings"
)

const (
	playerUrl  = Host + "/player/%s-%s-%s.html"
	viewUrl    = Host + "/view/%s.html"
	resolveUrl = "https://api.xiaomingming.org/cloud/mp6.php?vid=%s"
)

func GetPlayerApi(id string, sourceId string, epId string) string {
	return fmt.Sprintf(playerUrl, id, sourceId, epId)
}

func GetViewApi(id string) string {
	return fmt.Sprintf(viewUrl, id)
}

func GetResolveApi(url string) string {
	return fmt.Sprintf(resolveUrl, url)
}

func GetInfo(id string, sourceId string, epId string) model.ApiResponse {
	result := vhttp.DecodeString(httpc.Get(GetPlayerApi(id, sourceId, epId), nil).String(),
		"gb2312")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	eps := map[string][]string{}
	hrefExp := regexp.MustCompile("[0-9]+-[0-9]+-[0-9]+")
	doc.Find("div[class=movurls] > ul > li > a").Each(func(i int, selection *goquery.Selection) {
		hLink, b := selection.Attr("href")
		if b {
			idSL := strings.Split(hrefExp.FindString(hLink), "-")
			_, b := eps[idSL[1]]
			if !b {
				eps[idSL[1]] = make([]string, 0)
			}
			eps[idSL[1]] = append(eps[idSL[1]], GenerateUniqueId(idSL[0], idSL[1], idSL[2]))
		}
	})
	title, b := utils.SliceString(regexp.MustCompile("xTitle='(.*)'").FindString(result), 8, -1)
	if !b {
		return model.CreateEmptyApiResponseByStatus(e.EXTERNAL_API_ERROR)
	}
	playdataUrl, b := utils.SliceString(regexp.MustCompile("src=\"/playdata/(.*)\"").FindString(result), 5, -1)
	if !b {
		return model.CreateEmptyApiResponseByStatus(e.EXTERNAL_API_ERROR)
	}

	pic, _ := utils.SliceString(regexp.MustCompile("'bdPic':'[^,]*',").FindString(result), 9, -1)
	return model.CreateApiResponseByStatus(e.SUCCESS, map[string]interface{}{
		"title":       title,
		"pic":         pic,
		"playdataUrl": playdataUrl,
		"episodes":    eps,
	})
}

func GetPlayData(playdataUrl string) model.ApiResponse {
	result := httpc.Get(Host+playdataUrl, nil).String()
	videoList, b := utils.SliceString(strings.ReplaceAll(regexp.MustCompile("\\[(.*)\\],").FindString(result), "'", "\""), 0, -1)
	if !b {
		return model.CreateEmptyApiResponseByStatus(e.EXTERNAL_API_ERROR)
	}
	gresult := gjson.Get(videoList, "@this.#.1.0")
	urls := make([]string, 0)
	for _, value := range gresult.Array() {
		urls = append(urls, strings.Split(value.String(), "$")[1])
	}
	return model.CreateApiResponseByStatus(e.SUCCESS, map[string]interface{}{
		"urls": urls,
	})
}

func ResolveVideoUrl(url string) model.ApiResponse {
	result := strings.ReplaceAll(httpc.Get(GetResolveApi(url), nil).String(), " ", "")
	realUrl, b := utils.SliceString(regexp.MustCompile("varvideo='(.*)';").FindString(result), 10, -2)
	if !b {
		return model.CreateEmptyApiResponseByStatus(e.EXTERNAL_API_ERROR)
	}
	return model.CreateApiResponseByStatus(e.SUCCESS, map[string]interface{}{
		"url":     url,
		"realUrl": realUrl,
	})
}
