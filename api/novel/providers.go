package novel

import (
	"AynaAPI/api/novel/core"
	"AynaAPI/api/novel/provider"
)

var ProviderMap map[string]core.NovelProvider

func init() {
	ProviderMap = map[string]core.NovelProvider{
		//provider.SobiqugeAPI.GetName():provider.SobiqugeAPI,
		provider.LiqugeAPI.GetName(): provider.LiqugeAPI,
		provider.BiquwxAPI.GetName(): provider.BiquwxAPI,
	}
}

func GetNovelProviderList() []string {
	plist := make([]string, 0)
	for key, _ := range ProviderMap {
		plist = append(plist, key)
	}
	return plist
}

func IsProviderAvailable(identifier string) bool {
	val, _ := ProviderMap[identifier]
	return val != nil
}

func GetNovelProvider(identifier string) core.NovelProvider {
	val, _ := ProviderMap[identifier]
	return val
}
