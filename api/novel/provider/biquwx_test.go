package provider

import (
	"AynaAPI/api/novel/core"
	"fmt"
	"testing"
)

func TestBiquwx_Search(t *testing.T) {
	var provider core.NovelProvider = BiquwxAPI
	fmt.Println(provider.Search("诡秘之主"))
}

func TestBiquwx_Novel(t *testing.T) {
	var provider core.NovelProvider = BiquwxAPI
	result, _ := provider.Search("诡秘之主")
	meta := result.Result[0]
	novel, _ := provider.GetNovel(meta)
	chapter := novel.Volumes[0].Chapters[0]
	fmt.Println(provider.UpdateNovelChapter(&chapter))
	fmt.Println(chapter)
}
