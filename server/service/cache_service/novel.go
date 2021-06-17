package cache_service

import (
	"fmt"
	"regexp"
)

const (
	PREFIX_NOVEL_INFO    = "api_novel_info"
	PREFIX_NOVEL_CONTENT = "api_novel_content"
	PREFIX_NOVEL_SEARCH  = "api_novel_search"
)

func formatUrl(uri string) string {
	return regexp.MustCompile("http(s)?://").ReplaceAllString(uri, "")
}

func GetNovelInfoKey(uri string) string {
	return PREFIX_PROVIDER_INFO + "_" + formatUrl(uri)
}

func GetNovelContentKey(uri string) string {
	return PREFIX_PROVIDER_PLAYURLS + "_" + formatUrl(uri)
}

func GetNovelSearchKey(providerName string, keyword string) string {
	return PREFIX_PROVIDER_SEARCH + "_" + fmt.Sprintf("%s-%s", providerName, keyword)
}
