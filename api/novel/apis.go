package novel

//
//import (
//	"AynaAPI/api/core"
//	e2 "AynaAPI/api/core/e"
//	"regexp"
//)
//
//func GetDataByProvider(provider *NovelProvider, uri string) core.ApiResponse {
//	if regexp.MustCompile(provider.InfoUrl).FindString(uri) != "" {
//		return GetInfoByProvider(provider, uri)
//	}
//	if regexp.MustCompile(provider.ContentUrl).FindString(uri) != "" {
//		return GetContentByProvider(provider, uri)
//	}
//	return core.CreateEmptyApiResponseByStatus(e2.NOVEL_PROVIDER_URL_NOT_SUPPORT)
//}
//
//func GetData(uri string) core.ApiResponse {
//	for _, provider := range ProviderMap {
//		result := GetDataByProvider(&provider, uri)
//		if result.Status != e2.NOVEL_PROVIDER_URL_NOT_SUPPORT {
//			return result
//		}
//	}
//	return core.CreateEmptyApiResponseByStatus(e2.NOVEL_PROVIDER_URL_NOT_SUPPORT_2)
//}
