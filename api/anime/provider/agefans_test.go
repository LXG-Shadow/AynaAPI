package provider

import (
	"AynaAPI/api/anime/core"
	"fmt"
	"testing"
)

func TestAgefans_Search(t *testing.T) {
	var provider core.AnimeProvider = AgefansAPI
	search, err := provider.Search("刀剑神域")
	if err != nil {
		return
	}
	fmt.Println(search)
}

func TestAgefans_GetAnimeMeta(t *testing.T) {
	var provider core.AnimeProvider = AgefansAPI
	fmt.Println(provider.GetAnimeMeta(core.ProviderMeta{
		Name: "agefans",
		Url:  "https://www.agefans.cc/detail/20190087",
	}))
}

func TestAgefans_GetAnime(t *testing.T) {
	var provider core.AnimeProvider = AgefansAPI
	search, _ := provider.Search("ggo")
	am := search.Result[0]
	fmt.Println(am)
	anime, _ := provider.GetAnime(am)
	fmt.Println(anime)
}

func TestAgefans_updateAnimeVideo(t *testing.T) {
	provider := AgefansAPI
	video := core.AnimeVideo{
		Title: "miao",
		Url:   "",
		Provider: core.ProviderMeta{
			Name: "",
			Url:  "https://www.agefans.cc/_getplay?aid=20210249&playindex=2&epindex=1",
		},
	}
	fmt.Println(provider.UpdateAnimeVideo(&video))
	fmt.Println(video)
}
