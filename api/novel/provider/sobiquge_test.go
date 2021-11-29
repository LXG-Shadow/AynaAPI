package provider

import (
	"AynaAPI/api/novel"
	"fmt"
	"testing"
)

func TestSobiquge_Search(t *testing.T) {
	var provider novel.NovelProvider = SobiqugeAPI
	fmt.Println(provider.Search("诡秘之主"))
}
