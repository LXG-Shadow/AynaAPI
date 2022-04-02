package resp

import (
	"AynaAPI/api/music"
	"AynaAPI/server/app"
)

type MusicInfo struct {
	app.AppJsonResponse
	Data music.Music
}

type MusicPlayUrl struct {
	app.AppJsonResponse
	Data music.MusicAudio
}

// MusicSearchResult todo: additionalProp1, no solution yet
type MusicSearchResult struct {
	app.AppJsonResponse
	Data map[string][]music.MusicMeta
}

type MusicProviderList struct {
	app.AppJsonResponse
	Data []string
}
