package provider

import (
	"AynaAPI/api/anime/core"
	"fmt"
	"testing"
)

func TestSusuDm_Search(t *testing.T) {
	var provider core.AnimeProvider = SusuDmAPI
	search, err := provider.Search("刀剑神域")
	if err != nil {
		return
	}
	for _, meta := range search.Result {
		//fmt.Println(meta)
		_ = provider.UpdateAnimeMeta(&meta)
		anime, err := provider.GetAnime(meta)
		fmt.Println(err)
		fmt.Println(anime)
		break
	}
}
