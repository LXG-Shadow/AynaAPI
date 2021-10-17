package provider

import (
	"AynaAPI/api/novel/core"
	"fmt"
	"testing"
)

func TestSobiquge_Search(t *testing.T) {
	var provider core.NovelProvider = SobiqugeAPI
	fmt.Println(provider.Search("诡秘之主"))
}
