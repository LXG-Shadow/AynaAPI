package provider

import (
	"AynaAPI/api/core"
	e2 "AynaAPI/api/e"
	"AynaAPI/api/httpc"
	"AynaAPI/api/novel"
	"AynaAPI/api/novel/rule"
	"AynaAPI/utils/vhttp"
	"fmt"
	"github.com/aynakeya/deepcolor"
	"regexp"
)

type Biquge struct {
	Name      string
	BaseUrl   string
	Charset   string
	SearchAPI string
	Header    map[string]string
	Rules     rule.BiqugeRule
}

var SobiqugeAPI *Biquge

func __newSobiquge() *Biquge {
	return &Biquge{
		Name:      "sobiquge",
		BaseUrl:   "https://www.sobiquge.com/",
		Charset:   "utf-8",
		SearchAPI: "https://www.sobiquge.com/search.php?q=%s",
		Rules:     rule.InitializeSobiqugeRules(),
	}
}

func init() {
	SobiqugeAPI = __newSobiquge()
}

func (p *Biquge) GetName() string {
	return p.Name
}

func (p *Biquge) Validate(meta core.ProviderMeta) bool {
	return meta.Name == p.GetName() &&
		regexp.MustCompile("^"+regexp.QuoteMeta(p.BaseUrl)).FindString(meta.Url) != ""
}

func (p *Biquge) GetNovelMeta(meta core.ProviderMeta) (novel.NovelMeta, error) {
	nMeta := novel.NovelMeta{Provider: meta}
	err := p.UpdateNovelMeta(&nMeta)
	return nMeta, err
}

func (p *Biquge) UpdateNovelMeta(meta *novel.NovelMeta) error {
	meta.Provider.Name = p.Name
	if meta.Provider.Url == "" {
		return e2.NewError(e2.INTERNAL_ERROR)
	}
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         meta.Provider.Url,
		Charset:     p.Charset,
		ContentType: deepcolor.TentacleContentTypeHTMl,
	}, httpc.GetCORSString, nil, nil)
	if err != nil {
		return err
	}
	meta.Title = result.GetSingle(p.Rules.Title)
	meta.Cover = result.GetSingle(p.Rules.Cover)
	meta.Author = result.GetSingle(p.Rules.Author)
	meta.Description = result.GetSingle(p.Rules.Description)
	return nil
}

func (p *Biquge) GetNovel(meta novel.NovelMeta) (novel.Novel, error) {
	novell := novel.Novel{NovelMeta: meta}
	err := p.UpdateNovel(&novell)
	return novell, err
}

func (p *Biquge) UpdateNovel(novell *novel.Novel) error {
	if novell.Provider.Url == "" {
		return e2.NewError(e2.INTERNAL_ERROR)
	}
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         novell.Provider.Url,
		Charset:     p.Charset,
		ContentType: deepcolor.TentacleContentTypeHTMl,
	}, httpc.GetCORSString, nil, nil)
	if err != nil {
		return err
	}
	novell.Volumes = make([]novel.NovelVolume, 0)
	volume := novel.NovelVolume{
		Title:    "正文",
		Chapters: make([]novel.NovelChapter, 0),
	}
	for _, chapter := range result.GetMapList(p.Rules.Chapters) {
		volume.Chapters = append(volume.Chapters, novel.NovelChapter{
			Title:   chapter["name"],
			Content: "",
			Provider: core.ProviderMeta{
				Name: p.Name,
				Url:  vhttp.CompleteUrl(novell.Provider.Url, chapter["url"]),
			},
		})
	}
	novell.Volumes = append(novell.Volumes, volume)
	return nil
}

func (p *Biquge) UpdateNovelChapter(chapter *novel.NovelChapter) error {
	if chapter.Provider.Url == "" {
		return e2.NewError(e2.INTERNAL_ERROR)
	}
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         chapter.Provider.Url,
		Charset:     p.Charset,
		ContentType: deepcolor.TentacleContentTypeHTMl,
	}, httpc.GetCORSString, nil, nil)
	if err != nil {
		return err
	}
	value := result.GetMap(p.Rules.Content)
	chapter.Content = value["content"]
	return nil
}

func (p *Biquge) Search(keyword string) (novel.NovelSearchResult, error) {
	uri := fmt.Sprintf(p.SearchAPI, keyword)
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         uri,
		Charset:     p.Charset,
		ContentType: deepcolor.TentacleContentTypeHTMl,
	}, httpc.GetCORSString, nil, nil)
	if err != nil {
		return novel.NovelSearchResult{}, err
	}
	searchResult := novel.NovelSearchResult{
		Result: make([]novel.NovelMeta, 0),
	}
	for _, sr := range result.GetMapList(p.Rules.Search) {
		searchResult.Result = append(searchResult.Result, novel.NovelMeta{
			Title:       sr["title"],
			Author:      sr["author"],
			Cover:       sr["cover"],
			Description: sr["description"],
			Provider: core.ProviderMeta{
				Name: p.Name,
				Url:  vhttp.CompleteUrl(p.BaseUrl, sr["url"]),
			},
		})
	}
	return searchResult, nil
}
