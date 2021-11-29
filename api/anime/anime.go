package anime

import (
	"AynaAPI/api/core"
	"encoding/json"
)

type AnimeMeta struct {
	Title       string            `json:"title"`
	Cover       string            `json:"cover"`
	Year        string            `json:"year"`
	Tags        []string          `json:"tags"`
	Description string            `json:"description"`
	Provider    core.ProviderMeta `json:"provider"`
}

func (a *AnimeMeta) GetCompletionStatus() bool {
	return a.Title != "" && a.Provider.GetCompletionStatus()
}

// MarshalJSON method from http://choly.ca/post/go-json-marshalling/
func (self *AnimeMeta) MarshalJSON() ([]byte, error) {
	type FakeV AnimeMeta
	return json.Marshal(&struct {
		Metadata string `json:"metadata"`
		*FakeV
	}{
		Metadata: self.Provider.Dump(),
		FakeV:    (*FakeV)(self),
	})
}

type Anime struct {
	AnimeMeta
	Playlists []Playlist `json:"playlists"`
}

func (a *Anime) GetCompletionStatus() bool {
	return a.AnimeMeta.GetCompletionStatus() && a.Playlists != nil
}

type Playlist struct {
	Name   string       `json:"name"`
	Videos []AnimeVideo `json:"videos"`
}

type AnimeVideo struct {
	Title    string            `json:"title"`
	Url      string            `json:"url"`
	Provider core.ProviderMeta `json:"provider"`
}

func (a *AnimeVideo) GetCompletionStatus() bool {
	return a.Url != ""
}
