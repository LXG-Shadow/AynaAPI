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

type AnimeProvider interface {
	GetName() string
	Validate(meta ProviderMeta) bool
	Search(keyword string) (AnimeSearchResult, error)
	GetAnimeMeta(meta ProviderMeta) (AnimeMeta, error)
	UpdateAnimeMeta(meta *AnimeMeta) error
	GetAnime(meta AnimeMeta) (Anime, error)
	UpdateAnime(anime *Anime) error
	UpdateAnimeVideo(video *AnimeVideo) error
}
