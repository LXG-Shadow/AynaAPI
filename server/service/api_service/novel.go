package api_service

import (
	apiModel "AynaAPI/api/model"
	novelApi "AynaAPI/api/novel"
	"AynaAPI/pkg/gredis"
	"AynaAPI/server/service/cache_service"
	"time"
)

func NovelSearch(provider *novelApi.NovelProvider, keyword string, useCache bool) apiModel.ApiResponse {
	if !gredis.Online {
		return novelApi.SearchByProvider(provider, keyword)
	}
	key := cache_service.GetNovelSearchKey(provider.Identifier, keyword)
	if !useCache {
		resp := novelApi.SearchByProvider(provider, keyword)
		defer gredis.SetData(key, resp, time.Hour*24)
		return resp
	}
	// redis has cache
	var resp apiModel.ApiResponse
	if b := gredis.GetData(key, &resp); b {
		return resp
	}
	resp = novelApi.SearchByProvider(provider, keyword)
	defer gredis.SetData(key, resp, time.Hour*24)
	return resp
}

func NovelContent(provider *novelApi.NovelProvider, uri string, useCache bool) apiModel.ApiResponse {
	if !gredis.Online {
		return novelApi.GetContentByProvider(provider, uri)
	}
	key := cache_service.GetNovelContentKey(uri)
	if !useCache {
		resp := novelApi.GetContentByProvider(provider, uri)
		defer gredis.SetData(key, resp, time.Hour*24*31)
		return resp
	}
	// redis has cache
	var resp apiModel.ApiResponse
	if b := gredis.GetData(key, &resp); b {
		return resp
	}
	resp = novelApi.GetContentByProvider(provider, uri)
	defer gredis.SetData(key, resp, time.Hour*24*31)
	return resp
}

func NovelInfo(provider *novelApi.NovelProvider, uri string, useCache bool) apiModel.ApiResponse {
	if !gredis.Online {
		return novelApi.GetInfoByProvider(provider, uri)
	}
	key := cache_service.GetNovelInfoKey(uri)
	if !useCache {
		resp := novelApi.GetInfoByProvider(provider, uri)
		defer gredis.SetData(key, resp, time.Hour*24*16)
		return resp
	}
	// redis has cache
	var resp apiModel.ApiResponse
	if b := gredis.GetData(key, &resp); b {
		return resp
	}
	resp = novelApi.GetInfoByProvider(provider, uri)
	defer gredis.SetData(key, resp, time.Hour*24*16)
	return resp
}
