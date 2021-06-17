package novel

import (
	"AynaAPI/api/httpc"
	"AynaAPI/api/model"
	"AynaAPI/api/model/e"
	"AynaAPI/pkg/deepcolor"
	"AynaAPI/utils/vhttp"
	"fmt"
)

func GetInfoByProvider(provider *NovelProvider, uri string) model.ApiResponse {
	result := httpc.GetCORS(uri).String()
	if result == "" {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	doc, err := deepcolor.NewDocumentFromStringWithTagRepl(result)
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	chapters := deepcolor.ParseMapList(doc, provider.Rule.Chapters)
	for _, chapter := range chapters {
		chaUrl := chapter["url"]
		if chaUrl != "" && !vhttp.IsUrl(chaUrl) {
			chapter["url"] = vhttp.JoinUrl(uri, chaUrl)
		}
	}
	return model.CreateApiResponseByStatus(e.SUCCESS, map[string]interface{}{
		"title":       deepcolor.ParseSingle(doc, provider.Rule.Title),
		"author":      deepcolor.ParseSingle(doc, provider.Rule.Author),
		"abstraction": deepcolor.ParseSingle(doc, provider.Rule.Abstraction),
		"cover":       deepcolor.ParseSingle(doc, provider.Rule.Cover),
		"chapters":    chapters,
	})
}

func GetContentByProvider(provider *NovelProvider, uri string) model.ApiResponse {
	result := httpc.GetCORS(uri).String()
	if result == "" {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	doc, err := deepcolor.NewDocumentFromStringWithTagRepl(result)
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	value := deepcolor.ParseMap(doc, provider.Rule.Content)
	return model.CreateApiResponseByStatus(e.SUCCESS, map[string]interface{}{
		"content": value["content"],
		"name":    value["name"],
	})
}

func SearchByProvider(provider *NovelProvider, keyword string) model.ApiResponse {
	uri := fmt.Sprintf(provider.SearchApi, keyword)
	result := httpc.GetCORS(uri).String()
	if result == "" {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	doc, err := deepcolor.NewDocumentFromStringWithTagRepl(result)
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	searchResult := deepcolor.ParseMapList(doc, provider.Rule.Search)
	for _, sr := range searchResult {
		bUrl := sr["url"]
		if bUrl != "" && !vhttp.IsUrl(bUrl) {
			sr["url"] = vhttp.JoinUrl(provider.HomeUrl, bUrl)
		}
	}
	return model.CreateApiResponseByStatus(e.SUCCESS, map[string]interface{}{
		"result": searchResult,
	})
}
