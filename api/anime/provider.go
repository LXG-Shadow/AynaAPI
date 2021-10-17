package anime

import (
	"AynaAPI/api/anime/core"
	"AynaAPI/api/anime/provider"
)

var ProviderMap map[string]core.AnimeProvider

func init() {
	ProviderMap = map[string]core.AnimeProvider{
		provider.AgefansAPI.GetName(): provider.AgefansAPI,
		provider.SusuDmAPI.GetName():  provider.SusuDmAPI,
	}
}

func GetAnimeProviderList() []string {
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

func GetAnimeProvider(identifier string) core.AnimeProvider {
	val, _ := ProviderMap[identifier]
	return val
}
