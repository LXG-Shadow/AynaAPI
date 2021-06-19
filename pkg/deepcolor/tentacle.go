package deepcolor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

type TentacleContentType string

const (
	TentacleContentTypeHTMl = "html"
	TentacleContentTypeJson = "json"
)

type Tentacle struct {
	Url         string              `json:"url"`
	Charset     string              `json:"charset"`
	ContentType TentacleContentType `json:"content_type"`
	Header      map[string]string   `json:"header"`
}

func TentacleHTML(uri, charset string) Tentacle {
	return Tentacle{
		Url:         uri,
		Charset:     charset,
		ContentType: TentacleContentTypeHTMl,
	}
}

func TentacleJson(uri, charset string) Tentacle {
	return Tentacle{
		Url:         uri,
		Charset:     charset,
		ContentType: TentacleContentTypeJson,
	}
}

type TentacleResult interface {
	GetSingle(item Item) string
	GetList(item Item) []string
	GetMap(item Item) map[string]string
	GetMapList(item Item) []map[string]string
}

type TentacleHTMLResult struct {
	Document *goquery.Document
}

func (t TentacleHTMLResult) GetSingle(item Item) string {
	return ParseSingle(t.Document, item)
}

func (t TentacleHTMLResult) GetList(item Item) []string {
	return ParseList(t.Document, item)
}

func (t TentacleHTMLResult) GetMap(item Item) map[string]string {
	return ParseMap(t.Document, item)
}

func (t TentacleHTMLResult) GetMapList(item Item) []map[string]string {
	return ParseMapList(t.Document, item)
}

type TentacleJsonResult struct {
	Document *gjson.Result
}

func (t TentacleJsonResult) GetSingle(item Item) string {
	return ParseJsonSingle(t.Document, item)
}

func (t TentacleJsonResult) GetList(item Item) []string {
	return ParseJsonList(t.Document, item)
}

func (t TentacleJsonResult) GetMap(item Item) map[string]string {
	return ParseJsonMap(t.Document, item)
}

func (t TentacleJsonResult) GetMapList(item Item) []map[string]string {
	return ParseJsonMapList(t.Document, item)
}
