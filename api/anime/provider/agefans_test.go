package provider

import (
	"AynaAPI/api/anime/core"
	"fmt"
	"testing"
)

func TestAgefans_Search(t *testing.T) {
	provider := AgefansAPI
	search, err := provider.Search("刀剑神域")
	if err != nil {
		return
	}
	fmt.Println(search)
}

func TestAgefans_GetAnimeMeta(t *testing.T) {
	var provider = AgefansAPI
	fmt.Println(provider.GetAnimeMeta(core.ProviderMeta{
		Name: "agefans",
		Url:  "https://www.agefans.cc/detail/20190087",
	}))
}
