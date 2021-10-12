package api_service

import (
	"AynaAPI/api/core"
	providerApi "AynaAPI/api/provider"
	"AynaAPI/pkg/gredis"
	"AynaAPI/server/service/cache_service"
	"time"
)

func ProviderInitialize(provider string, model *providerApi.ApiProvider, useCache bool) bool {
	if !gredis.Online {
		return (*model).Initialize()
	}
	key := cache_service.GetProviderInfoKey(model)
	if !useCache {
		defer gredis.SetData(key, model, time.Hour*24)
		return (*model).Initialize()
	}
	// redis has cache
	if s, b := gredis.GetString(key); b {
		if new_model, ok := providerApi.InitWithJsonData(provider, []byte(s)); ok {
			(*model) = *new_model
			return true
		}
	}
	defer gredis.SetData(key, model, time.Hour*24)
	return (*model).Initialize()
}

func ProviderGetPlayUrls(model *providerApi.ApiProvider, useCache bool) []core.ApiResource {
	if !gredis.Online {
		return (*model).GetPlayUrls()
	}
	key := cache_service.GetProviderPlayUrlsKey(model)
	if !useCache {
		defer gredis.SetData(key, model, time.Hour*24*3)
		return (*model).GetPlayUrls()
	}
	// redis has cache
	var resource []core.ApiResource
	if b := gredis.GetData(key, &resource); b {
		return resource
	}
	defer gredis.SetData(key, model, time.Hour*24*3)
	return (*model).GetPlayUrls()
}

func ProviderSearch(provider string, keyword string, page int, useCache bool) core.ApiResponse {
	if !gredis.Online {
		return providerApi.Search(provider, keyword, page)
	}
	key := cache_service.GetProviderSearchKey(provider, keyword, page)
	if !useCache {
		result := providerApi.Search(provider, keyword, page)
		gredis.SetData(key, result, time.Hour*24)
		return result
	}
	var result core.ApiResponse
	if b := gredis.GetData(key, &result); b {
		return result
	}
	result = providerApi.Search(provider, keyword, page)
	gredis.SetData(key, result, time.Hour*24)
	return result
}
