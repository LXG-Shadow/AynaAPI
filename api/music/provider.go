package music

import (
	"AynaAPI/api/core"
)

type MusicProvider interface {
	GetName() string
	Validate(meta core.ProviderMeta) bool
	Search(keyword string) (MusicSearchResult, error)
	GetMusicMeta(meta core.ProviderMeta) (MusicMeta, error)
	UpdateMusicMeta(meta *MusicMeta) error
	GetMusic(meta MusicMeta) (Music, error)
	UpdateMusic(anime *Music) error
	UpdateMusicAudio(audio *MusicAudio) error
}

type MusicProviderManager struct {
	core.ProviderManager
}

func (p *MusicProviderManager) GetProvider(identifier string) MusicProvider {
	val, _ := p.ProviderMap[identifier]
	if val == nil {
		return nil
	}
	return val.(MusicProvider)
}

var Providers *MusicProviderManager

func init() {
	Providers = &MusicProviderManager{core.ProviderManager{ProviderMap: map[string]interface{}{}}}
}
