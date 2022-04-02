package provider

import (
	"AynaAPI/api/anime"
	"AynaAPI/api/core"
	"fmt"
	"testing"
)

func TestAgefans_Search(t *testing.T) {
	var provider anime.AnimeProvider = AgefansAPI
	search, err := provider.Search("刀剑神域")
	if err != nil {
		return
	}
	fmt.Println(search)
}

func TestAgefans_GetAnimeMeta(t *testing.T) {
	var provider anime.AnimeProvider = AgefansAPI
	fmt.Println(provider.GetAnimeMeta(core.ProviderMeta{
		Name: "agefans",
		Url:  "20190087",
	}))
}

func TestAgefans_GetAnime(t *testing.T) {
	var provider anime.AnimeProvider = AgefansAPI
	search, _ := provider.Search("ggo")
	am := search.Result[0]
	//fmt.Println(am)
	animee, _ := provider.GetAnime(am)
	fmt.Println(animee.Playlists)
}

func TestAgefans_updateAnimeVideo(t *testing.T) {
	provider := AgefansAPI
	video := anime.AnimeVideo{
		Title: "miao",
		Url:   "",
		Provider: core.ProviderMeta{
			Name: "",
			Url:  "20210249-2-1",
		},
	}
	fmt.Println(provider.UpdateAnimeVideo(&video))
	fmt.Println(video)
}
