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
	"regexp"
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

func (p *SusuDm) getSearchApi(keyword string) string {
	return fmt.Sprintf(p.SearchAPI, keyword, 1)
}

func (p *SusuDm) GetAnimeMeta(meta core.ProviderMeta) core.AnimeMeta {
	meta.Name = "susudm"
	aMeta := core.AnimeMeta{
		Provider: meta,
	}
	p.UpdateAnimeMeta(&aMeta)
	return aMeta
}

func (p *SusuDm) UpdateAnimeMeta(meta *core.AnimeMeta) apiCore.ApiResponse {
	if regexp.MustCompile("/[0-9]+/").FindString(meta.Provider.Url) == "" {
		return apiCore.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         meta.Provider.Url,
		Charset:     "utf-8",
		ContentType: deepcolor.TentacleContentTypeHTMl,
	}, httpc.GetCORS, nil, nil)
	if err != nil {
		return apiCore.CreateEmptyApiResponseByStatus(e.EXTERNAL_API_ERROR)
	}
	meta.Title = strings.TrimSpace(result.GetSingle(p.Rules.Title))
	meta.Cover = result.GetSingle(p.Rules.Cover)
	meta.Tags = result.GetList(p.Rules.Tags)
	meta.Description = result.GetSingle(p.Rules.Desc)
	tmp := strings.Split(result.GetSingle(p.Rules.AreaYear), "---")
	if len(tmp) < 2 {
		tmp = strings.Split(result.GetSingle(p.Rules.AreaYear2), "---")
	}
	if len(tmp) < 2 {
		meta.Year = "-1"
	} else {
		meta.Year = tmp[1]
	}
	return apiCore.CreateEmptyApiResponseByStatus(e.SUCCESS)
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
	var sResults = make([]core.AnimeMeta, 0)
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
	return apiCore.CreateEmptyApiResponseByStatus(e.EXTERNAL_API_ERROR)
}
