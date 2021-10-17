package api_service

import (
	novelCore "AynaAPI/api/novel/core"
	"AynaAPI/pkg/gredis"
	"fmt"
	"testing"
)

func TestNovelSearch(t *testing.T) {
	gredis.Initialize()
	fmt.Println(NovelSearch("liquge", "诡秘之主", true))
	fmt.Println(NovelSearch("biquwx", "诡秘之主", true))
}

func TestNovelGet(t *testing.T) {
	gredis.Initialize()
	meta1 := novelCore.ProviderMeta{
		Name: "liquge",
		Url:  "http://www.liquge.com/book/18641/",
	}
	meta2 := novelCore.ProviderMeta{
		Name: "biquwx",
		Url:  "https://www.biquwx.la/3_3746/",
	}
	fmt.Println(NovelGet(meta1.Dump(), true))
	fmt.Println(NovelGet(meta2.Dump(), true))
}

func TestNovelContent(t *testing.T) {
	gredis.Initialize()
	meta1 := novelCore.ProviderMeta{
		Name: "liquge",
		Url:  "http://www.liquge.com/book/18641/",
	}
	meta2 := novelCore.ProviderMeta{
		Name: "biquwx",
		Url:  "https://www.biquwx.la/3_3746/",
	}
	fmt.Println(NovelContent(meta1.Dump(), 0, 0, true))
	fmt.Println(NovelContent(meta2.Dump(), 0, 0, true))
}
