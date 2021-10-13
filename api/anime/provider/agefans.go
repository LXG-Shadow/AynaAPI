package provider

import (
	"AynaAPI/api/anime/core"
	"AynaAPI/api/anime/rule"
	"AynaAPI/api/core/e"
	"AynaAPI/api/httpc"
	"AynaAPI/utils/vhttp"
	"fmt"
	"github.com/aynakeya/deepcolor"
	"math"
	"regexp"
	"strings"
	"time"
)

type Agefans struct {
	BaseUrl    string
	SearchAPI  string
	PlayUrlAPI string
	Rules      rule.AgefansRules
}

func _newAgefans() *Agefans {
	return &Agefans{
		BaseUrl:    "https://www.agefans.cc",
		SearchAPI:  "https://www.agefans.cc/search?query=%s&page=%d",
		PlayUrlAPI: "http://d.gqyy8.com:8077/ne2/s%d.js?%d",
		Rules:      rule.InitializeAgefansRules(),
	}
}

var AgefansAPI *Agefans

func init() {
	AgefansAPI = _newAgefans()
}

func (p *Agefans) getSearchApi(keyword string) string {
	return fmt.Sprintf(p.SearchAPI, keyword, 1)
}

func (p *Agefans) GetAnimeMeta(meta core.ProviderMeta) (core.AnimeMeta, error) {
	aMeta := core.AnimeMeta{Provider: meta}
	err := p.UpdateAnimeMeta(&aMeta)
	return aMeta, err
}

func (p *Agefans) UpdateAnimeMeta(meta *core.AnimeMeta) error {
	id := regexp.MustCompile("/detail/[0-9]+").FindString(meta.Provider.Url)
	if id == "" {
		return e.NewError(e.INTERNAL_ERROR)
	}
	meta.Provider.Name = "agefans"
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         meta.Provider.Url,
		Charset:     "utf-8",
		ContentType: deepcolor.TentacleContentTypeHTMl,
	}, httpc.GetCORS, nil, nil)
	if err != nil {
		return e.NewError(e.EXTERNAL_API_ERROR)
	}
	meta.Title = result.GetSingle(p.Rules.InfoTitle)
	meta.Year = result.GetSingle(p.Rules.InfoYear)
	meta.Cover = result.GetSingle(p.Rules.InfoCover)
	meta.Description = result.GetSingle(p.Rules.InfoDesc)
	meta.Tags = strings.Split(result.GetSingle(p.Rules.InfoTag), " ")
	return nil
}

//func (p *Agefans) GetAnime(meta core.AnimeMeta) (core.Anime, error){
//
//}
//func (p *Agefans) UpdateAnime(anime *core.Anime) error{
//
//}

func (p *Agefans) Search(keyword string) (core.AnimeSearchResult, error) {
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         p.getSearchApi(keyword),
		Charset:     "utf-8",
		ContentType: deepcolor.TentacleContentTypeHTMl,
	}, httpc.GetCORS, nil, nil)

	if err != nil {
		return core.AnimeSearchResult{}, e.NewError(e.EXTERNAL_API_ERROR)
	}
	var sResults = make([]core.AnimeMeta, 0)
	urls := result.GetList(p.Rules.SearchURL)
	titles := result.GetList(p.Rules.SearchTitle)
	years := result.GetList(p.Rules.SearchYear)
	tags := result.GetList(p.Rules.SearchTag)
	covers := result.GetList(p.Rules.SearchCover)
	desc := result.GetList(p.Rules.SearchDesc)
	for index, url := range urls {
		meta := core.AnimeMeta{
			Title:       titles[index],
			Year:        years[index],
			Tags:        strings.Split(tags[index], " "),
			Cover:       covers[index],
			Description: desc[index],
			Provider: core.ProviderMeta{
				Name: "agefans",
				Url:  vhttp.JoinUrl(p.BaseUrl, url),
			},
		}
		sResults = append(sResults, meta)
	}
	return core.AnimeSearchResult{Result: sResults}, nil
}

func (p *Agefans) getCookie(t1 int) string {
	timeNow := time.Now().UnixNano() / (1000000)
	t1Tmp := int64(math.Round(float64(t1)/1000)) >> 0x5
	k2 := (t1Tmp*(t1Tmp%0x1000)*0x3+0x1450f)*(t1Tmp%0x1000) + t1Tmp
	t2 := timeNow
	t2 = t2 - t2%10 + k2%10
	return fmt.Sprintf("t1=%d;k2=%d;t2=%d", t1, k2, t2)
}
