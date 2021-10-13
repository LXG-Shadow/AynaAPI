package provider

import (
	"fmt"
	"testing"
)

func TestSusuDm_Search(t *testing.T) {
	provider := _newSusuDm()
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
