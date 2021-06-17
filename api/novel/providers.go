package novel

import (
	"regexp"
)

var ProviderMap map[string]NovelProvider = map[string]NovelProvider{
	BiqugeProvider.Identifier:  BiqugeProvider,
	BiqugeBProvider.Identifier: BiqugeBProvider,
	BiqugeCProvider.Identifier: BiqugeCProvider,
}

func IsProviderAvailable(identifier string) bool {
	p, ok := ProviderMap[identifier]
	if !ok {
		return false
	}
	return p.Status
}

func IsNovelProviderExists(identifier string) bool {
	_, ok := ProviderMap[identifier]
	return ok
}

func GetNovelProvider(identifier string) *NovelProvider {
	val, _ := ProviderMap[identifier]
	return &val
}

func GetNovelProviderByUrl(uri string) *NovelProvider {
	for _, provider := range ProviderMap {
		if regexp.MustCompile(provider.InfoUrl).FindString(uri) != "" {
			return &provider
		}
		if regexp.MustCompile(provider.ContentUrl).FindString(uri) != "" {
			return &provider
		}
	}
	return nil
}
