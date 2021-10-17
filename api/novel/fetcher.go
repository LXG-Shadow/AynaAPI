package novel

//
//import (
//	"AynaAPI/api/core"
//	e2 "AynaAPI/api/core/e"
//	"AynaAPI/api/httpc"
//	"AynaAPI/pkg/deepcolor"
//	"AynaAPI/utils/vhttp"
//	"fmt"
//)
//
//func completeUrl(host, path string) string {
//	if path == "" {
//		return ""
//	}
//	if !vhttp.IsUrl(path) {
//		return vhttp.JoinUrl(host, path)
//	}
//	return path
//}
//
//func GetInfoByProvider(provider *NovelProvider, uri string) core.ApiResponse {
//	result := httpc.GetCORS(uri).String()
//	if result == "" {
//		return core.CreateEmptyApiResponseByStatus(e2.INTERNAL_ERROR)
//	}
//	doc, err := deepcolor.NewDocumentFromStringWithEncoding(result, provider.Charset)
//	if err != nil {
//		return core.CreateEmptyApiResponseByStatus(e2.INTERNAL_ERROR)
//	}
//	chapaterUrl := completeUrl(uri, deepcolor.ParseSingle(doc, provider.Rule.ChapaterUrl))
//	var chapters = []map[string]string{}
//	if chapaterUrl == "" {
//		chapters = deepcolor.ParseMapList(doc, provider.Rule.Chapters)
//	} else {
//		for chapaterUrl != "" {
//			rs, err1 := deepcolor.Fetch(deepcolor.TentacleHTML(chapaterUrl, provider.Charset), deepcolor.GetCORS)
//			if err1 == nil && rs != nil {
//				chapters = append(chapters, rs.GetMapList(provider.Rule.Chapters)...)
//				chapaterUrl = completeUrl(uri, rs.GetSingle(provider.Rule.ChapaterUrl))
//			} else {
//				break
//			}
//
//		}
//	}
//	for _, chapter := range chapters {
//		chapter["url"] = completeUrl(uri, chapter["url"])
//	}
//	return core.CreateApiResponseByStatus(e2.SUCCESS, map[string]interface{}{
//		"title":       deepcolor.ParseSingle(doc, provider.Rule.Title),
//		"author":      deepcolor.ParseSingle(doc, provider.Rule.Author),
//		"abstraction": deepcolor.ParseSingle(doc, provider.Rule.Abstraction),
//		"cover":       deepcolor.ParseSingle(doc, provider.Rule.Cover),
//		"chapters":    chapters,
//	})
//}
//
//func GetContentByProvider(provider *NovelProvider, uri string) core.ApiResponse {
//	result, err := deepcolor.Fetch(deepcolor.Tentacle{
//		Url:         uri,
//		Charset:     provider.Charset,
//		ContentType: deepcolor.TentacleContentTypeHTMl,
//		Header:      provider.Header,
//	}, deepcolor.GetCORS)
//	if err != nil {
//		return core.CreateEmptyApiResponseByStatus(e2.INTERNAL_ERROR)
//	}
//	value := result.GetMap(provider.Rule.Content)
//	contentUrl := completeUrl(uri, result.GetSingle(provider.Rule.ContentUrl))
//	for contentUrl != "" {
//		rs, err1 := deepcolor.Fetch(deepcolor.Tentacle{
//			Url:         contentUrl,
//			Charset:     provider.Charset,
//			ContentType: deepcolor.TentacleContentTypeHTMl,
//			Header:      provider.Header,
//		}, deepcolor.GetCORS)
//		if err1 == nil {
//			value["content"] = value["content"] + rs.GetMap(provider.Rule.Content)["content"]
//			contentUrl = completeUrl(uri, rs.GetSingle(provider.Rule.ContentUrl))
//		} else {
//			break
//		}
//	}
//	return core.CreateApiResponseByStatus(e2.SUCCESS, map[string]interface{}{
//		"content": value["content"],
//		"name":    value["name"],
//		"info":    completeUrl(uri, value["info"]),
//	})
//}
//
//func SearchByProvider(provider *NovelProvider, keyword string) core.ApiResponse {
//	uri := fmt.Sprintf(provider.SearchApi, keyword)
//	result, err := deepcolor.Fetch(deepcolor.Tentacle{
//		Url:         uri,
//		Charset:     provider.Charset,
//		ContentType: deepcolor.TentacleContentTypeHTMl,
//		Header:      provider.Header,
//	}, deepcolor.GetCORS)
//	if err != nil {
//		return core.CreateEmptyApiResponseByStatus(e2.INTERNAL_ERROR)
//	}
//	searchResult := result.GetMapList(provider.Rule.Search)
//	for _, sr := range searchResult {
//		bUrl := sr["url"]
//		if bUrl != "" && !vhttp.IsUrl(bUrl) {
//			sr["url"] = vhttp.JoinUrl(provider.HomeUrl, bUrl)
//		}
//	}
//	return core.CreateApiResponseByStatus(e2.SUCCESS, map[string]interface{}{
//		"result": searchResult,
//	})
//}
