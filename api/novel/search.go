package novel

import (
	"AynaAPI/api/core"
)

type NovelSearchResult struct {
	core.SearchResult
	Result []NovelMeta `json:"result"`
}
