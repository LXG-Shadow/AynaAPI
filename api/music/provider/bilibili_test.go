package provider

import (
	"AynaAPI/api/core"
	"AynaAPI/api/music"
	"fmt"
	"testing"
)

func TestBilibili_Search(t *testing.T) {
	var api music.MusicProvider = BiliMusicAPI
	rs, _ := api.Search("霜雪千年")
	fmt.Println(rs)
	fmt.Println(api.Validate(rs.Result[0].Provider))
	fmt.Println(api.Search("asdbfdadfas"))
}

func TestBilibili_GetMusicMeta(t *testing.T) {
	var api music.MusicProvider = BiliMusicAPI
	meta := core.ProviderMeta{
		Name: api.GetName(),
		Url:  "au79970",
	}
	musicMeta, err := api.GetMusicMeta(meta)
	if err != nil {
		return
	}
	fmt.Println(musicMeta)
}

func TestBilibili_GetMusic(t *testing.T) {
	var api music.MusicProvider = BiliMusicAPI
	meta := core.ProviderMeta{
		Name: api.GetName(),
		Url:  "au79970",
	}
	musicMeta, err := api.GetMusicMeta(meta)
	if err != nil {
		return
	}
	musicc, err := api.GetMusic(musicMeta)
	if err != nil {
		return
	}
	fmt.Println(musicc)
	api.UpdateMusicAudio(&musicc.Audio)
	fmt.Println(musicc.Audio)
}
