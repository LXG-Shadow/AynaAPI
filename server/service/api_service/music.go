package api_service

import (
	"AynaAPI/api/core"
	"AynaAPI/api/music"
	_ "AynaAPI/api/music/provider"
	"AynaAPI/pkg/gredis"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/cache_service"
)

func MusicGetAudio(metadata string, ua string, useCache bool) (music.MusicAudio, int) {
	musicc, errcode := MusicGet(metadata, useCache)
	if errcode != 0 {
		return musicc.Audio, errcode
	}
	audio := musicc.Audio
	audio.UserAgent = ua
	err := music.Providers.GetProvider(musicc.Provider.Name).UpdateMusicAudio(&audio)
	if err != nil {
		return audio, e.MUSIC_GET_DATA_FAIL
	}
	return audio, 0
}

func MusicGet(metadata string, useCache bool) (music.Music, int) {
	meta := core.ProviderMeta{}
	err := meta.Load(metadata)
	if err != nil {
		return music.Music{}, e.MUSIC_INITIALIZE_FAIL
	}
	key := cache_service.GetMusicKey(meta)

	if useCache && gredis.Online {
		var result music.Music
		if b := gredis.GetData(key, &result); b {
			return result, 0
		}
	}

	result, errcode := _MusicGet(meta)
	if errcode != 0 {
		return result, errcode
	}
	if gredis.Online {
		defer gredis.SetData(key, result, cache_service.GetCacheExpirePeriod())
	}
	return result, 0
}

func _MusicGet(meta core.ProviderMeta) (musicc music.Music, errcode int) {
	musicc = music.Music{}
	errcode = 0
	var provider music.MusicProvider
	for _, providerName := range music.Providers.GetProviderList() {
		if music.Providers.GetProvider(providerName).Validate(meta) {
			provider = music.Providers.GetProvider(providerName)
			break
		}
	}
	if provider == nil {
		errcode = e.MUSIC_PROVIDER_NOT_AVAILABLE
		return
	}
	aMeta, err := provider.GetMusicMeta(meta)
	if err != nil {
		errcode = e.MUSIC_INITIALIZE_FAIL
		return
	}
	musicc, err = provider.GetMusic(aMeta)
	if err != nil {
		errcode = e.MUSIC_INITIALIZE_FAIL
		return
	}
	return
}

func MusicSearch(providerName string, keyword string, useCache bool) (music.MusicSearchResult, int) {
	provider := music.Providers.GetProvider(providerName)
	if provider == nil {
		return music.MusicSearchResult{}, e.MUSIC_PROVIDER_NOT_AVAILABLE
	}
	key := cache_service.GetMusicSearchKey(provider, keyword)

	if useCache && gredis.Online {
		var result music.MusicSearchResult
		if b := gredis.GetData(key, &result); b {
			return result, 0
		}
	}
	result, err := provider.Search(keyword)
	if err != nil {
		return result, e.MUSIC_SEARCH_FAIL
	}
	if gredis.Online {
		defer gredis.SetData(key, result, cache_service.GetCacheExpirePeriod())
	}
	return result, 0
}
