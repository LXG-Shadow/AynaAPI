package deepcolor

import (
	"errors"
	"github.com/tidwall/gjson"
)

type RequestFunc func(uri string, header map[string]string) string
type ProcessFunc func(result TentacleResult) bool

func executeHandlers(tentacleResult TentacleResult, handlers []ProcessFunc) {
	for _, processFunc := range handlers {
		if !processFunc(tentacleResult) {
			return
		}
	}
}

func Fetch(tentacle Tentacle, requestFunc RequestFunc, handlers ...ProcessFunc) (TentacleResult, error) {
	switch tentacle.ContentType {
	case TentacleContentTypeHTMl:
		return FetchHTML(tentacle, requestFunc, handlers...)
	case TentacleContentTypeJson:
		return FetchJson(tentacle, requestFunc, handlers...)
	default:
		return FetchHTML(tentacle, requestFunc, handlers...)
	}
}

func FetchHTML(tentacle Tentacle, requestFunc RequestFunc, handlers ...ProcessFunc) (TentacleResult, error) {
	result := requestFunc(tentacle.Url, tentacle.Header)
	if result == "" {
		return nil, errors.New("http connection error")
	}
	doc, err := NewDocumentFromStringWithEncoding(result, tentacle.Charset)
	if err != nil {
		return nil, err
	}
	tentacleResult := TentacleHTMLResult{Document: doc}
	defer executeHandlers(tentacleResult, handlers)
	return tentacleResult, nil
}

func FetchJson(tentacle Tentacle, requestFunc RequestFunc, handlers ...ProcessFunc) (TentacleResult, error) {
	result := requestFunc(tentacle.Url, tentacle.Header)
	if result == "" {
		return nil, errors.New("http connection error")
	}
	jsonResult := gjson.Parse(result)
	tentacleResult := TentacleJsonResult{Document: &jsonResult}
	defer executeHandlers(tentacleResult, handlers)
	return tentacleResult, nil
}
