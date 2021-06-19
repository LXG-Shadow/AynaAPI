package novel

import (
	"AynaAPI/api/httpc"
	"AynaAPI/api/model"
	"AynaAPI/api/model/e"
	"AynaAPI/pkg/deepcolor"
	"AynaAPI/utils/vhttp"
	"fmt"
)

func completeUrl(host, path string) string {
	if path == "" {
		return ""
	}
	if !vhttp.IsUrl(path) {
		return vhttp.JoinUrl(host, path)
	}
	return path
}

func GetInfoByProvider(provider *NovelProvider, uri string) model.ApiResponse {
	result := httpc.GetCORS(uri).String()
	if result == "" {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	doc, err := deepcolor.NewDocumentFromStringWithEncoding(result, provider.Charset)
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	chapaterUrl := completeUrl(uri, deepcolor.ParseSingle(doc, provider.Rule.ChapaterUrl))
	var chapters = []map[string]string{}
	if chapaterUrl == "" {
		chapters = deepcolor.ParseMapList(doc, provider.Rule.Chapters)
	} else {
		for chapaterUrl != "" {
			rs, err1 := deepcolor.Fetch(deepcolor.TentacleHTML(chapaterUrl, provider.Charset), deepcolor.GetCORS)
			if err1 == nil && rs != nil {
				chapters = append(chapters, rs.GetMapList(provider.Rule.Chapters)...)
				chapaterUrl = completeUrl(uri, rs.GetSingle(provider.Rule.ChapaterUrl))
			} else {
				break
			}

		}
	}
	for _, chapter := range chapters {
		chapter["url"] = completeUrl(uri, chapter["url"])
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
	doc, err := deepcolor.NewDocumentFromStringWithEncoding(result, provider.Charset)
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	value := deepcolor.ParseMap(doc, provider.Rule.Content)
	contentUrl := completeUrl(uri, deepcolor.ParseSingle(doc, provider.Rule.ContentUrl))
	for contentUrl != "" {
		rs, err1 := deepcolor.Fetch(deepcolor.TentacleHTML(contentUrl, provider.Charset), deepcolor.GetCORS)
		if err1 == nil {
			value["content"] = value["content"] + rs.GetMap(provider.Rule.Content)["content"]
			contentUrl = completeUrl(uri, rs.GetSingle(provider.Rule.ContentUrl))
		} else {
			break
		}
	}
	return model.CreateApiResponseByStatus(e.SUCCESS, map[string]interface{}{
		"content": value["content"],
		"name":    value["name"],
	})
}

func SearchByProvider(provider *NovelProvider, keyword string) model.ApiResponse {
	uri := fmt.Sprintf(provider.SearchApi, keyword)
	result, err := deepcolor.Fetch(deepcolor.Tentacle{
		Url:         uri,
		Charset:     provider.Charset,
		ContentType: deepcolor.TentacleContentTypeHTMl,
		Header:      provider.Header,
	}, deepcolor.GetCORS)
	if err != nil {
		return model.CreateEmptyApiResponseByStatus(e.INTERNAL_ERROR)
	}
	searchResult := result.GetMapList(provider.Rule.Search)
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
