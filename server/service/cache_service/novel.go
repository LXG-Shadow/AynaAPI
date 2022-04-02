package cache_service

import (
	"AynaAPI/api/core"
	"AynaAPI/api/novel"
	"fmt"
)

const (
	PREFIX_NOVEL_INFO   = "api_novel_info"
	PREFIX_NOVEL_SEARCH = "api_novel_search"
)

func GetNovelInfoKey(meta core.ProviderMeta) string {
	return PREFIX_NOVEL_INFO + "_" + meta.Dump()
}

func GetNovelSearchKey(provider novel.NovelProvider, keyword string) string {
	return PREFIX_NOVEL_SEARCH + "_" + fmt.Sprintf("%s_%s", provider.GetName(), keyword)
}
