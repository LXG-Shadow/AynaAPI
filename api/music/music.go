package music

import (
	"AynaAPI/api/core"
	"encoding/json"
)

type MusicMeta struct {
	Title    string            `json:"title"`
	Cover    string            `json:"cover"`
	Author   string            `json:"author"`
	Provider core.ProviderMeta `json:"provider"`
}

func (a *MusicMeta) GetCompletionStatus() bool {
	return a.Title != "" && a.Provider.GetCompletionStatus()
}

// MarshalJSON method from http://choly.ca/post/go-json-marshalling/
func (self *MusicMeta) MarshalJSON() ([]byte, error) {
	type FakeV MusicMeta
	return json.Marshal(&struct {
		Metadata string `json:"metadata"`
		*FakeV
	}{
		Metadata: self.Provider.Dump(),
		FakeV:    (*FakeV)(self),
	})
}

type Music struct {
	MusicMeta
	Audio MusicAudio `json:"audio"`
	Lyric MusicLyric `json:"lyric"`
}

type MusicAudio struct {
	Url       string            `json:"url"`
	UserAgent string            `json:"user_agent"`
	Provider  core.ProviderMeta `json:"provider"`
}

func (a *MusicAudio) GetCompletionStatus() bool {
	return a.Url != ""
}

type MusicLyric struct {
	Data string `json:"data"`
}
