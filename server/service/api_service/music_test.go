package api_service

import (
	"fmt"
	"testing"
)

func TestMusicSearch(t *testing.T) {
	fmt.Println(MusicSearch("bilibili", "刀剑神域", false))
	fmt.Println(MusicSearch("bilibilimusic", "霜雪千年", false))
	fmt.Println(MusicSearch("bilibilimusic", "霜雪千年", false))
}
