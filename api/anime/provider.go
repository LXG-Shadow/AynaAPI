package anime

import (
	"AynaAPI/api/core"
)

type AnimeProvider interface {
	GetName() string
	Validate(meta core.ProviderMeta) bool
	Search(keyword string) (AnimeSearchResult, error)
	GetAnimeMeta(meta core.ProviderMeta) (AnimeMeta, error)
	UpdateAnimeMeta(meta *AnimeMeta) error
	GetAnime(meta AnimeMeta) (Anime, error)
	UpdateAnime(anime *Anime) error
	UpdateAnimeVideo(video *AnimeVideo) error
}

type AnimeProviderManager struct {
	core.ProviderManager
}

func (p *AnimeProviderManager) GetProvider(identifier string) AnimeProvider {
	val, _ := p.ProviderMap[identifier]
	if val == nil {
		return nil
	}
	return val.(AnimeProvider)
}

var Providers *AnimeProviderManager

func init() {
	Providers = &AnimeProviderManager{core.ProviderManager{ProviderMap: map[string]interface{}{}}}
}
