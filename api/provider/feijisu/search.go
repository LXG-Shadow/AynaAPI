package susudm

import (
	"AynaAPI/api/core"
	e2 "AynaAPI/api/core/e"
	"AynaAPI/api/httpc"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

const (
	searchApi = "http://159.75.7.49:7211/sssv.php?q=%s&top=%s"
)

func GetSearchApi(keyword string, page int) string {
	return fmt.Sprintf(searchApi, keyword, page*10)
}

type searchResult struct {
	Url       string `json:"url"`
	Thumb     string `json:"thumb"`
	Title     string `json:"title"`
	Time      string `json:"time"`
	Catid     string `json:"catid"`
	Star      string `json:"star"`
	Lianzaijs string `json:"lianzaijs"`
	Beizhu    string `json:"beizhu"`
	AliasFull string `json:"alias_full"`
	Area      string `json:"area"`
	Sort      string `json:"sort"`
}

func Search(keyword string, page int) core.ApiResponse {
	if page == 0 {
		page = 1
	}
	result := httpc.Get(GetSearchApi(keyword, page), map[string]string{"origin": Host}).String()
	if result == "" {
		return core.CreateEmptyApiResponseByStatus(e2.INTERNAL_ERROR)
	}
	videoList := make([]*FeijisuDmVideo, 0)
	var sresults []searchResult
	err := json.Unmarshal([]byte(gjson.Parse(strings.ReplaceAll(result, "\ufeff", "")).String()), &sresults)
	if err != nil {
		return core.CreateEmptyApiResponseByStatus(e2.EXTERNAL_API_ERROR)
	}
	for _, rs := range sresults {
		if v, b := InitWithUrl(rs.Url); b {
			v.PictureUrl = rs.Thumb
			v.Title = rs.Title
			videoList = append(videoList, v)
		}
	}
	return core.CreateApiResponseByStatus(e2.SUCCESS, map[string]interface{}{
		"videoList":   videoList,
		"currentPage": page,
		"totalPage":   page,
	})
}
