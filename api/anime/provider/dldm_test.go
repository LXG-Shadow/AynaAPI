package provider

import (
	"AynaAPI/api/anime"
	"AynaAPI/api/core"
	"fmt"
	"testing"
)

func TestDldm_Search(t *testing.T) {
	var provider anime.AnimeProvider = DldmAPI
	search, err := provider.Search("刀剑神域")
	if err != nil {
		return
	}
	fmt.Println(search)
}

func TestDldm_GetAnimeMeta(t *testing.T) {
	var provider anime.AnimeProvider = DldmAPI
	m, e := provider.GetAnimeMeta(core.ProviderMeta{
		Name: "dldm",
		Url:  "2240",
	})
	fmt.Println(m)
	fmt.Println(e)
}

func TestDldm_GetAnime(t *testing.T) {
	var provider anime.AnimeProvider = DldmAPI
	search, _ := provider.Search("刀剑神域")
	am := search.Result[0]
	animee, _ := provider.GetAnime(am)
	fmt.Println(animee.Playlists)
}

func TestDldm_updateAnimeVideo(t *testing.T) {
	var provider anime.AnimeProvider = DldmAPI
	video := anime.AnimeVideo{
		Title: "miao",
		Url:   "",
		Provider: core.ProviderMeta{
			Name: "",
			Url:  "444-1-1",
			//Url: "16536-1-1",
		},
	}
	fmt.Println(provider.UpdateAnimeVideo(&video))
	fmt.Println(video)
}
