package novel

import (
	"AynaAPI/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
)

var ProviderMap map[string]NovelProvider = map[string]NovelProvider{
	//BiqugeProvider.Identifier:  BiqugeProvider,
	//BiqugeBProvider.Identifier: BiqugeBProvider,
	//BiqugeCProvider.Identifier: BiqugeCProvider,
	LigntNovelProvider.Identifier: LigntNovelProvider,
}

func init() {
	var count int = 0
	fs, _ := ioutil.ReadDir(config.APIConfig.NovelRulePath)
	for _, file := range fs {
		if !file.IsDir() {
			content, err := ioutil.ReadFile(path.Join(config.APIConfig.NovelRulePath, file.Name()))
			if err != nil {
				continue
			}
			var tmp []NovelProvider
			if err := json.Unmarshal(content, &tmp); err != nil {
				fmt.Println(err)
				continue
			}
			for _, pv := range tmp {
				ProviderMap[pv.Identifier] = pv
				count++
			}
		}
	}
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
