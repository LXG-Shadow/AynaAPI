package provider

import (
	"AynaAPI/api/anime"
	"AynaAPI/api/core"
	"fmt"
	"testing"
)

func TestOmofun_Search(t *testing.T) {
	var provider anime.AnimeProvider = OmofunAPI
	search, err := provider.Search("刀剑神域")
	if err != nil {
		return
	}
	fmt.Println(search)
}

func TestOmofun_GetAnimeMeta(t *testing.T) {
	var provider anime.AnimeProvider = OmofunAPI
	fmt.Println(provider.GetAnimeMeta(core.ProviderMeta{
		Name: "omofun",
		Url:  "2240",
	}))
}

func TestAOmofun_GetAnime(t *testing.T) {
	var provider anime.AnimeProvider = OmofunAPI
	search, _ := provider.Search("天才王子的赤字国家振兴术")
	am := search.Result[0]
	animee, _ := provider.GetAnime(am)
	fmt.Println(animee.Playlists)
}

func TestOmofun_updateAnimeVideo(t *testing.T) {
	var provider anime.AnimeProvider = OmofunAPI
	video := anime.AnimeVideo{
		Title: "miao",
		Url:   "",
		Provider: core.ProviderMeta{
			Name: "",
			Url:  "5277-1-1",
		},
	}
	fmt.Println(provider.UpdateAnimeVideo(&video))
	fmt.Println(video)
}
