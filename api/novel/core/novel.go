package core

import "encoding/json"

type NovelMeta struct {
	Title       string       `json:"title"`
	Author      string       `json:"cover"`
	Cover       string       `json:"year"`
	Description string       `json:"description"`
	Provider    ProviderMeta `json:"provider"`
}

func (a *NovelMeta) GetCompletionStatus() bool {
	return a.Title != "" && a.Provider.GetCompletionStatus()
}

// MarshalJSON method from http://choly.ca/post/go-json-marshalling/
func (self *NovelMeta) MarshalJSON() ([]byte, error) {
	type FakeV NovelMeta
	return json.Marshal(&struct {
		Metadata string `json:"metadata"`
		*FakeV
	}{
		Metadata: self.Provider.Dump(),
		FakeV:    (*FakeV)(self),
	})
}

type Novel struct {
	NovelMeta
	Volumes []NovelVolume `json:"volumes"`
}

func (a *Novel) GetCompletionStatus() bool {
	return a.NovelMeta.GetCompletionStatus() && a.Volumes != nil
}

type NovelVolume struct {
	Title    string         `json:"title"`
	Chapters []NovelChapter `json:"chapters"`
}

type NovelChapter struct {
	Title    string       `json:"title"`
	Content  string       `json:"content"`
	Provider ProviderMeta `json:"provider"`
}

func (a *NovelChapter) GetCompletionStatus() bool {
	return a.Content != ""
}
