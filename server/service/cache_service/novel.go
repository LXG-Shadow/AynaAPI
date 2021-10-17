package cache_service

import (
	novelCore "AynaAPI/api/novel/core"
	"fmt"
)

const (
	PREFIX_NOVEL_INFO   = "api_novel_info"
	PREFIX_NOVEL_SEARCH = "api_novel_search"
)

func GetNovelInfoKey(meta novelCore.ProviderMeta) string {
	return PREFIX_NOVEL_INFO + "_" + meta.Dump()
}

func GetNovelSearchKey(provider novelCore.NovelProvider, keyword string) string {
	return PREFIX_NOVEL_SEARCH + "_" + fmt.Sprintf("%s_%s", provider.GetName(), keyword)
}
