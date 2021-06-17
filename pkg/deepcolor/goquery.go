package deepcolor

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

var (
	TagReplacement = map[string]string{"<br(\\s)?(/)?>": "\n"}
)

func NewDocumentFromStringWithTagRepl(src string) (*goquery.Document, error) {
	for tag, repl := range TagReplacement {
		src = regexp.MustCompile(tag).ReplaceAllString(src, repl)
	}
	return goquery.NewDocumentFromReader(strings.NewReader(src))
}
