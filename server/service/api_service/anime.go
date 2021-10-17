package api_service

import (
	"AynaAPI/api/anime"
	animeCore "AynaAPI/api/anime/core"
	"AynaAPI/pkg/gredis"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/cache_service"
	"time"
)

func AnimeGetVideo(metadata string, playlist int, episode int, useCache bool) (animeCore.AnimeVideo, int) {
	animee, errcode := AnimeGet(metadata, useCache)
	if errcode != 0 {
		return animeCore.AnimeVideo{}, errcode
	}
	if playlist < 0 || playlist >= len(animee.Playlists) {
		return animeCore.AnimeVideo{}, e.BGM_VIDEO_NOT_FOUND
	}
	if episode < 0 || episode >= len(animee.Playlists[playlist].Videos) {
		return animeCore.AnimeVideo{}, e.BGM_VIDEO_NOT_FOUND
	}
	v := &animee.Playlists[playlist].Videos[episode]
	if v.GetCompletionStatus() && useCache {
		return *v, 0
	}
	err := anime.GetAnimeProvider(animee.Provider.Name).UpdateAnimeVideo(v)
	if err != nil {
		return *v, e.BGM_VIDEO_NOT_FOUND
	}
	key := cache_service.GetAnimeKey(animee.Provider)
	if gredis.Online {
		defer gredis.SetData(key, animee, time.Hour*24)
	}
	return *v, 0
}

func AnimeGet(metadata string, useCache bool) (animeCore.Anime, int) {
	meta := animeCore.ProviderMeta{}
	err := meta.Load(metadata)
	if err != nil {
		return animeCore.Anime{}, e.BGM_INITIALIZE_FAIL
	}
	key := cache_service.GetAnimeKey(meta)

	if useCache && gredis.Online {
		var result animeCore.Anime
		if b := gredis.GetData(key, &result); b {
			return result, 0
		}
	}

	result, errcode := _AnimeGet(meta)
	if errcode != 0 {
		return result, errcode
	}
	if gredis.Online {
		defer gredis.SetData(key, result, time.Hour*24)
	}
	return result, 0
}

func _AnimeGet(meta animeCore.ProviderMeta) (animee animeCore.Anime, errcode int) {
	animee = animeCore.Anime{}
	errcode = 0
	var provider animeCore.AnimeProvider
	for _, providerName := range anime.GetAnimeProviderList() {
		if anime.GetAnimeProvider(providerName).Validate(meta) {
			provider = anime.GetAnimeProvider(providerName)
			break
		}
	}
	if provider == nil {
		errcode = e.BGM_PROVIDER_NOT_AVAILABLE
		return
	}
	aMeta, err := provider.GetAnimeMeta(meta)
	if err != nil {
		errcode = e.BGM_INITIALIZE_FAIL
		return
	}
	animee, err = provider.GetAnime(aMeta)
	if err != nil {
		errcode = e.BGM_INITIALIZE_FAIL
		return
	}
	return
}

func AnimeSearch(providerName string, keyword string, useCache bool) (animeCore.AnimeSearchResult, int) {
	provider := anime.GetAnimeProvider(providerName)
	if provider == nil {
		return animeCore.AnimeSearchResult{}, e.BGM_PROVIDER_NOT_AVAILABLE
	}
	key := cache_service.GetAnimeSearchKey(provider, keyword)

	if useCache && gredis.Online {
		var result animeCore.AnimeSearchResult
		if b := gredis.GetData(key, &result); b {
			return result, 0
		}
	}
	result, err := provider.Search(keyword)
	if err != nil {
		return result, e.BGM_SEARCH_FAIL
	}
	if gredis.Online {
		defer gredis.SetData(key, result, time.Hour*24)
	}
	return result, 0
}
