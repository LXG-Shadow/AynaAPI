package core

import "AynaAPI/api/core"

type ProviderMeta struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Provider interface {
	Search(keyword string) core.ApiResource
	GetAnimeMeta()
	UpdateAnimeMeta(meta *AnimeMeta)
}
