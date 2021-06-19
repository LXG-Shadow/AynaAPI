package novel

import (
	"AynaAPI/pkg/deepcolor"
	"regexp"
)

type NovelProvider struct {
	Identifier string            `json:"identifier"`
	Name       string            `json:"name"`
	Alias      string            `json:"alias"`
	HomeUrl    string            `json:"home_url"`
	Charset    string            `json:"charset"`
	Header     map[string]string `json:"header"`

	Status bool `json:"status"`

	InfoUrl    string `json:"info_url"`
	ContentUrl string `json:"content_url"`

	SearchApi string `json:"search_api"`

	Rule NovelProviderRule `json:"rule"`
}

func (self *NovelProvider) IsInfoUrl(uri string) bool {
	return regexp.MustCompile(self.InfoUrl).FindString(uri) != ""
}

func (self *NovelProvider) IsContentUrl(uri string) bool {
	return regexp.MustCompile(self.ContentUrl).FindString(uri) != ""
}

type NovelProviderRule struct {
	Title       deepcolor.Item `json:"title"`
	Author      deepcolor.Item `json:"author"`
	Cover       deepcolor.Item `json:"cover"`
	Abstraction deepcolor.Item `json:"abstraction"`

	Chapters    deepcolor.Item `json:"chapters"`
	ChapaterUrl deepcolor.Item `json:"chapater_url"`

	Content    deepcolor.Item `json:"content"`
	ContentUrl deepcolor.Item `json:"content_url"`

	Search deepcolor.Item `json:"search"`
}
