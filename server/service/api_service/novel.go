package api_service

import (
	"AynaAPI/api/novel"
	novelCore "AynaAPI/api/novel/core"
	"AynaAPI/pkg/gredis"
	"AynaAPI/server/app/e"
	"AynaAPI/server/service/cache_service"
	"time"
)

func NovelSearch(providerName string, keyword string, useCache bool) (novelCore.NovelSearchResult, int) {
	provider := novel.GetNovelProvider(providerName)
	if provider == nil {
		return novelCore.NovelSearchResult{}, e.NOVEL_PROVIDER_NOT_AVAILABLE
	}
	key := cache_service.GetNovelSearchKey(provider, keyword)

	if useCache && gredis.Online {
		var result novelCore.NovelSearchResult
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

func NovelContent(metadata string, volume int, chapter int, useCache bool) (novelCore.NovelChapter, int) {
	novell, errcode := NovelGet(metadata, useCache)
	if errcode != 0 {
		return novelCore.NovelChapter{}, errcode
	}
	if volume < 0 || volume >= len(novell.Volumes) {
		return novelCore.NovelChapter{}, e.NOVEL_CHAPTER_NOT_FOUND
	}
	if chapter < 0 || chapter >= len(novell.Volumes[volume].Chapters) {
		return novelCore.NovelChapter{}, e.NOVEL_CHAPTER_NOT_FOUND
	}
	c := &novell.Volumes[volume].Chapters[chapter]
	if c.GetCompletionStatus() && useCache {
		return *c, 0
	}
	err := novel.GetNovelProvider(novell.Provider.Name).UpdateNovelChapter(c)
	if err != nil {
		return *c, e.NOVEL_CHAPTER_NOT_FOUND
	}
	key := cache_service.GetNovelInfoKey(novell.Provider)
	if gredis.Online {
		defer gredis.SetData(key, novell, time.Hour*24)
	}
	return *c, 0
}

func NovelGet(metadata string, useCache bool) (novelCore.Novel, int) {
	meta := novelCore.ProviderMeta{}
	err := meta.Load(metadata)
	if err != nil {
		return novelCore.Novel{}, e.NOVEL_INITIALIZE_FAIL
	}
	key := cache_service.GetNovelInfoKey(meta)

	if useCache && gredis.Online {
		var result novelCore.Novel
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

func _NovelGet(meta novelCore.ProviderMeta) (novell novelCore.Novel, errcode int) {
	novell = novelCore.Novel{}
	errcode = 0
	var provider novelCore.NovelProvider
	for _, providerName := range novel.GetNovelProviderList() {
		if novel.GetNovelProvider(providerName).Validate(meta) {
			provider = novel.GetNovelProvider(providerName)
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
