package novel

import (
	"AynaAPI/api/core"
)

type NovelProvider interface {
	Search(keyword string) (NovelSearchResult, error)
	GetName() string
	Validate(meta core.ProviderMeta) bool
	GetNovelMeta(meta core.ProviderMeta) (NovelMeta, error)
	UpdateNovelMeta(meta *NovelMeta) error
	GetNovel(meta NovelMeta) (Novel, error)
	UpdateNovel(anime *Novel) error
	UpdateNovelChapter(video *NovelChapter) error
}

type NovelProviderManager struct {
	core.ProviderManager
}

func (p *NovelProviderManager) GetProvider(identifier string) NovelProvider {
	val, _ := p.ProviderMap[identifier]
	if val == nil {
		return nil
	}
	return val.(NovelProvider)
}

var Providers *NovelProviderManager

func init() {
	Providers = &NovelProviderManager{}
}
