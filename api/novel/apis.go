package novel

import (
	"AynaAPI/api/model"
	"AynaAPI/api/model/e"
	"regexp"
)

func GetDataByProvider(provider *NovelProvider, uri string) model.ApiResponse {
	if regexp.MustCompile(provider.InfoUrl).FindString(uri) != "" {
		return GetInfoByProvider(provider, uri)
	}
	if regexp.MustCompile(provider.ContentUrl).FindString(uri) != "" {
		return GetContentByProvider(provider, uri)
	}
	return model.CreateEmptyApiResponseByStatus(e.NOVEL_PROVIDER_URL_NOT_SUPPORT)
}

func GetData(uri string) model.ApiResponse {
	for _, provider := range ProviderMap {
		result := GetDataByProvider(&provider, uri)
		if result.Status != e.NOVEL_PROVIDER_URL_NOT_SUPPORT {
			return result
		}
	}
	return model.CreateEmptyApiResponseByStatus(e.NOVEL_PROVIDER_URL_NOT_SUPPORT_2)
}
