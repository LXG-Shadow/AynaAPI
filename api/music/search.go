package music

import "AynaAPI/api/core"

type MusicSearchResult struct {
	core.SearchResult
	Result []MusicMeta `json:"result"`
}
