package novel

import (
	deepcolor2 "AynaAPI/pkg/deepcolor"
	"regexp"
)

type NovelProvider struct {
	Identifier string
	Name       string
	Alias      string
	HomeUrl    string
	Charset    string

	Status bool

	InfoUrl    string
	ContentUrl string

	SearchApi string

	Rule NovelProviderRule
}

func (self *NovelProvider) IsInfoUrl(uri string) bool {
	return regexp.MustCompile(self.InfoUrl).FindString(uri) != ""
}

func (self *NovelProvider) IsContentUrl(uri string) bool {
	return regexp.MustCompile(self.ContentUrl).FindString(uri) != ""
}

type NovelProviderRule struct {
	Title       deepcolor2.RuleCollection
	Author      deepcolor2.RuleCollection
	Cover       deepcolor2.RuleCollection
	Abstraction deepcolor2.RuleCollection

	Chapters deepcolor2.RuleCollection

	Content deepcolor2.RuleCollection

	Search deepcolor2.RuleCollection
}
