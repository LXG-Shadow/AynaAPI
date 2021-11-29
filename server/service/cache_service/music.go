package cache_service

import (
	"AynaAPI/api/core"
	"AynaAPI/api/music"
	"fmt"
)

const (
	PREFIX_MUSIC_INFO   = "api_music_info"
	PREFIX_MUSIC_SEARCH = "api_music_search"
)

func GetMusicKey(meta core.ProviderMeta) string {
	return PREFIX_MUSIC_INFO + "_" + meta.Dump()
}

func GetMusicSearchKey(provider music.MusicProvider, keyword string) string {
	return PREFIX_MUSIC_SEARCH + "_" + fmt.Sprintf("%s_%s", provider.GetName(), keyword)
}
