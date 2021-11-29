package cache_service

import (
	"AynaAPI/api/anime"
	"AynaAPI/api/core"
	"fmt"
)

const (
	PREFIX_ANIME_INFO   = "api_anime_info"
	PREFIX_ANIME_SEARCH = "api_anime_search"
)

func GetAnimeKey(meta core.ProviderMeta) string {
	return PREFIX_ANIME_INFO + "_" + meta.Dump()
}

func GetAnimeSearchKey(provider anime.AnimeProvider, keyword string) string {
	return PREFIX_ANIME_SEARCH + "_" + fmt.Sprintf("%s_%s", provider.GetName(), keyword)
}
