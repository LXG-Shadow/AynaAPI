package core

import (
	"AynaAPI/utils/vhash"
	"fmt"
)

type AnimeMeta struct {
	Title       string       `json:"title"`
	Cover       string       `json:"cover"`
	Year        string       `json:"year"`
	Tags        []string     `json:"tags"`
	Description string       `json:"description"`
	Provider    ProviderMeta `json:"provider"`
}

type Anime struct {
	AnimeMeta
	Playlists []Playlist `json:"playlists"`
}

type Playlist struct {
	Name   string       `json:"name"`
	Videos []AnimeVideo `json:"videos"`
}

type AnimeVideo struct {
	Title    string       `json:"title"`
	Url      string       `json:"url"`
	Provider ProviderMeta `json:"provider"`
}

func (a *Anime) GetUID() string {
	return vhash.GetMD5Hash(fmt.Sprintf("%s-%s", a.Provider.Name, a.Title))
}
