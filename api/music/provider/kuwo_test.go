package provider

import (
	"AynaAPI/api/core"
	"AynaAPI/api/music"
	"fmt"
	"testing"
)

func TestKuwo_Search(t *testing.T) {
	var api music.MusicProvider = KuwoMusicAPI
	result, err := api.Search("霜雪千年")
	if err != nil {
		return
	}
	fmt.Println(result)
}

func TestKuwo_GetMusicMeta(t *testing.T) {
	var api music.MusicProvider = KuwoMusicAPI
	meta := core.ProviderMeta{
		Name: api.GetName(),
		Url:  "kuwo22804772",
	}
	musicMeta, err := api.GetMusicMeta(meta)
	fmt.Println(err)
	if err != nil {
		return
	}
	fmt.Println(musicMeta)
}

func TestKuwo_GetMusic(t *testing.T) {
	var api music.MusicProvider = KuwoMusicAPI
	meta := core.ProviderMeta{
		Name: api.GetName(),
		Url:  "kuwo22804772",
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
