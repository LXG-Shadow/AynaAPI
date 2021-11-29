package anime

import "AynaAPI/api/core"

type AnimeSearchResult struct {
	core.SearchResult
	Result []AnimeMeta `json:"result"`
}
