package core

type ProviderMeta struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Provider interface {
	Search(keyword string) (AnimeSearchResult, error)
	GetAnimeMeta(meta ProviderMeta) (AnimeMeta, error)
	UpdateAnimeMeta(meta *AnimeMeta) error
	GetAnime(meta AnimeMeta) (Anime, error)
	UpdateAnime(anime *Anime) error
}
