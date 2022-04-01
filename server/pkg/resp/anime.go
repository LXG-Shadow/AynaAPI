package resp

import (
	"AynaAPI/api/anime"
	"AynaAPI/server/app"
)

type AnimeInfo struct {
	app.AppJsonResponse
	Data anime.Anime
}

type AnimePlayUrl struct {
	app.AppJsonResponse
	Data anime.AnimeVideo
}

// AnimeSearchResult todo: additionalProp1, no solution yet
type AnimeSearchResult struct {
	app.AppJsonResponse
	Data map[string][]anime.AnimeMeta
}

type AnimeProviderList struct {
	app.AppJsonResponse
	Data []string
}
