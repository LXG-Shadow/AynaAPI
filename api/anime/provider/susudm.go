package provider

import (
	"AynaAPI/api/anime/core"
	"AynaAPI/api/anime/rule"
	apiCore "AynaAPI/api/core"
	"AynaAPI/api/core/e"
	"AynaAPI/api/httpc"
	"AynaAPI/utils/vhttp"
	"fmt"
	"github.com/aynakeya/deepcolor"
	"github.com/tidwall/gjson"
	"strings"
)

type SusuDm struct {
	BaseUrl   string
	SearchAPI string
	Rules     rule.SusuDmRules
}

func _newSusuDm() *SusuDm {
	return &SusuDm{
		BaseUrl:   "http://www.susudm.com",
		SearchAPI: "http://119.29.15.48:12334/ssszz.php?q=%s&top=%s",
		Rules:     rule.InitializeSusuDmRules(),
	}
}

func (p *SusuDm) getSearchApi(keywrod string) string {
	return fmt.Sprintf(p.SearchAPI, keywrod, 1)
}

func (p *SusuDm) GetAnimeMeta() {

}

func (p *SusuDm) UpdateAnimeMeta(meta *core.AnimeMeta) {

}

func (p *SusuDm) Search(keyword string) apiCore.ApiResponse {
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         p.getSearchApi(keyword),
		Charset:     "utf-8",
		ContentType: deepcolor.TentacleContentTypeText,
	}, httpc.GetCORS, nil, nil)
	if err != nil {
		return apiCore.CreateEmptyApiResponseByStatus(e.EXTERNAL_API_ERROR)
	}
	var sResults []core.AnimeMeta = make([]core.AnimeMeta, 0)
	jsonResult := gjson.Parse(
		strings.ReplaceAll(result.(deepcolor.TentacleTextResult).Data.(string), "\ufeff", ""))
	jsonResult.ForEach(func(key, value gjson.Result) bool {
		pMeta := core.ProviderMeta{
			Name: "susudm",
			Url:  vhttp.JoinUrl(p.BaseUrl, vhttp.GetUrlPath("http://"+value.Get("url").String())),
		}
		aMeta := core.AnimeMeta{
			Title:    value.Get("title").String(),
			Cover:    value.Get("thumb").String(),
			Year:     value.Get("time").String(),
			Provider: pMeta,
		}
		sResults = append(sResults, aMeta)
		return true
	})
	fmt.Println(sResults)
	return apiCore.CreateEmptyApiResponseByStatus(e.EXTERNAL_API_ERROR)
}
