package yhdm

import (
	"AynaAPI/api/httpc"
	"AynaAPI/api/model"
	"AynaAPI/api/model/e"
	"AynaAPI/utils/vhttp"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cast"
	"regexp"
	"strings"
)

const (
	searchApi = Host + "/search/%s/?page=%s"
)

func GetSearchApi(keyword string, page int) string {
	return fmt.Sprintf(searchApi, vhttp.QueryEscapeWithEncoding(keyword, "utf-8"), page)
}

func Search(keyword string, page int) model.ApiResponse {
	if page == 0 {
		page = 1
	}
	result := httpc.Get(GetSearchApi(keyword, page), nil).String()
	if result == "" {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	tP := doc.Find("a[id=lastn]").Text()
	totalPage, b := cast.ToIntE(tP)
	if b != nil {
		totalPage = page
	}
	videoList := make([]*YhdmVideo, 0)
	doc.Find("div[class~=fire] > div[class=lpic] > ul > li").Each(func(i int, selection *goquery.Selection) {
		info := selection.Find("h2 > a")
		if href, b := info.Attr("href"); b {
			title, _ := info.Attr("title")
			v := InitDefault()
			v.Id = regexp.MustCompile("[0-9]+").FindString(href)
			v.Title = title
			if pic, b := selection.Find("a > img").Attr("src"); b {
				v.PictureUrl = pic
			}
			videoList = append(videoList, v)
		}
	})
	return model.CreateApiResponseByStatus(e.SUCCESS, map[string]interface{}{
		"videoList":   videoList,
		"currentPage": page,
		"totalPage":   totalPage,
	})
}
