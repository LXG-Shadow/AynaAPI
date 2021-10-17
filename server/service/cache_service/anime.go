package cache_service

import (
	animeCore "AynaAPI/api/anime/core"
	"fmt"
)

const (
	PREFIX_ANIME_INFO   = "api_anime_info"
	PREFIX_ANIME_SEARCH = "api_anime_search"
)

func GetAnimeKey(meta animeCore.ProviderMeta) string {
	return PREFIX_ANIME_INFO + "_" + meta.Dump()
}

func GetAnimeSearchKey(provider animeCore.AnimeProvider, keyword string) string {
	return PREFIX_ANIME_SEARCH + "_" + fmt.Sprintf("%s_%s", provider.GetName(), keyword)
}
