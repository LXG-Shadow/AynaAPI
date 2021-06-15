package cache_service

import (
	"AynaAPI/api/provider"
	"fmt"
)

const (
	PREFIX_PROVIDER_INFO     = "api_provider_info"
	PREFIX_PROVIDER_PLAYURLS = "api_provider_playurls"
	PREFIX_PROVIDER_SEARCH   = "api_provider_search"
)

func GetProviderInfoKey(provider *provider.ApiProvider) string {
	return PREFIX_PROVIDER_INFO + "_" + (*provider).GetUniqueId()
}

func GetProviderPlayUrlsKey(provider *provider.ApiProvider) string {
	return PREFIX_PROVIDER_PLAYURLS + "_" + (*provider).GetUniqueId()
}

func GetProviderSearchKey(providerName string, keyword string, page int) string {
	return PREFIX_PROVIDER_SEARCH + "_" + fmt.Sprintf("%s-%s-%d", providerName, keyword, page)
}
