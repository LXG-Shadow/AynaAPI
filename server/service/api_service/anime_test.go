package api_service

import (
	animeCore "AynaAPI/api/anime/core"
	"AynaAPI/pkg/gredis"
	"fmt"
	"testing"
)

func TestAnimeSearch(t *testing.T) {
	gredis.Initialize()
	fmt.Println(AnimeSearch("susudm", "刀剑神域", false))
	fmt.Println(AnimeSearch("susudm", "刀剑神域", true))
	fmt.Println(AnimeSearch("agefans", "刀剑神域", true))
}

func TestAnimeGet(t *testing.T) {
	gredis.Initialize()
	meta1 := animeCore.ProviderMeta{
		Name: "susudm",
		Url:  "http://www.susudm.com/acg/38695/",
	}
	meta2 := animeCore.ProviderMeta{
		Name: "agefans",
		Url:  "https://www.agefans.cc/detail/20120038",
	}
	fmt.Println(AnimeGet(meta1.Dump(), true))
	fmt.Println(AnimeGet(meta2.Dump(), true))
}

func TestAnimeGetVideo(t *testing.T) {
	gredis.Initialize()
	meta1 := animeCore.ProviderMeta{
		Name: "susudm",
		Url:  "http://www.susudm.com/acg/38695/",
	}
	meta2 := animeCore.ProviderMeta{
		Name: "agefans",
		Url:  "https://www.agefans.cc/detail/20120038",
	}
	fmt.Println(AnimeGetVideo(meta1.Dump(), 0, 0, true))
	fmt.Println(AnimeGetVideo(meta2.Dump(), 0, 0, true))
}
