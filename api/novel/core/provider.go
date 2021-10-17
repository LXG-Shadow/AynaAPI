package core

import (
	"encoding/hex"
	"encoding/json"
)

type ProviderMeta struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (p *ProviderMeta) Load(data string) error {
	bdata, err := hex.DecodeString(data)
	if err != nil {
		return nil
	}
	return json.Unmarshal(bdata, p)
}

func (p *ProviderMeta) Dump() string {
	data, _ := json.Marshal(p)
	return hex.EncodeToString(data)
}

func (p *ProviderMeta) GetCompletionStatus() bool {
	return p.Url != ""
}

type NovelProvider interface {
	Search(keyword string) (NovelSearchResult, error)
	GetName() string
	Validate(meta ProviderMeta) bool
	GetNovelMeta(meta ProviderMeta) (NovelMeta, error)
	UpdateNovelMeta(meta *NovelMeta) error
	GetNovel(meta NovelMeta) (Novel, error)
	UpdateNovel(anime *Novel) error
	UpdateNovelChapter(video *NovelChapter) error
}
