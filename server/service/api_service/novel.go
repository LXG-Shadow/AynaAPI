package api_service

import (
	"AynaAPI/api/core"
	"AynaAPI/api/novel"
	"AynaAPI/pkg/gredis"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/cache_service"
	"time"

	// load all provider
	_ "AynaAPI/api/novel/provider"
)

func NovelSearch(providerName string, keyword string, useCache bool) (novel.NovelSearchResult, int) {
	provider := novel.Providers.GetProvider(providerName)
	if provider == nil {
		return novel.NovelSearchResult{}, e.NOVEL_PROVIDER_NOT_AVAILABLE
	}
	key := cache_service.GetNovelSearchKey(provider, keyword)

	if useCache && gredis.Online {
		var result novel.NovelSearchResult
		if b := gredis.GetData(key, &result); b {
			return result, 0
		}
	}
	result, err := provider.Search(keyword)
	if err != nil {
		return result, e.NOVEL_SEARCH_FAIL
	}
	if gredis.Online {
		defer gredis.SetData(key, result, time.Hour*24)
	}
	return result, 0
}

func NovelContent(metadata string, volume int, chapter int, useCache bool) (novel.NovelChapter, int) {
	novell, errcode := NovelGet(metadata, useCache)
	if errcode != 0 {
		return novel.NovelChapter{}, errcode
	}
	if volume < 0 || volume >= len(novell.Volumes) {
		return novel.NovelChapter{}, e.NOVEL_CHAPTER_NOT_FOUND
	}
	if chapter < 0 || chapter >= len(novell.Volumes[volume].Chapters) {
		return novel.NovelChapter{}, e.NOVEL_CHAPTER_NOT_FOUND
	}
	c := &novell.Volumes[volume].Chapters[chapter]
	if c.GetCompletionStatus() && useCache {
		return *c, 0
	}
	err := novel.Providers.GetProvider(novell.Provider.Name).UpdateNovelChapter(c)
	if err != nil {
		return *c, e.NOVEL_CHAPTER_NOT_FOUND
	}
	key := cache_service.GetNovelInfoKey(novell.Provider)
	if gredis.Online {
		defer gredis.SetData(key, novell, time.Hour*24)
	}
	return *c, 0
}

func NovelGet(metadata string, useCache bool) (novel.Novel, int) {
	meta := core.ProviderMeta{}
	err := meta.Load(metadata)
	if err != nil {
		return novel.Novel{}, e.NOVEL_INITIALIZE_FAIL
	}
	key := cache_service.GetNovelInfoKey(meta)

	if useCache && gredis.Online {
		var result novel.Novel
		if b := gredis.GetData(key, &result); b {
			return result, 0
		}
	}

	result, errcode := _NovelGet(meta)
	if errcode != 0 {
		return result, errcode
	}
	if gredis.Online {
		defer gredis.SetData(key, result, time.Hour*24)
	}
	return result, 0
}

func _NovelGet(meta core.ProviderMeta) (novell novel.Novel, errcode int) {
	novell = novel.Novel{}
	errcode = 0
	var provider novel.NovelProvider
	for _, providerName := range novel.Providers.GetProviderList() {
		if novel.Providers.GetProvider(providerName).Validate(meta) {
			provider = novel.Providers.GetProvider(providerName)
			break
		}
	}
	if provider == nil {
		errcode = e.NOVEL_PROVIDER_NOT_AVAILABLE
		return
	}
	aMeta, err := provider.GetNovelMeta(meta)
	if err != nil {
		errcode = e.NOVEL_INITIALIZE_FAIL
		return
	}
	novell, err = provider.GetNovel(aMeta)
	if err != nil {
		errcode = e.NOVEL_INITIALIZE_FAIL
		return
	}
	return
}
