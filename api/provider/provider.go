package provider

import (
	"AynaAPI/api/model"
	imomoeApi "AynaAPI/api/provider/imomoe"
	yhdmApi "AynaAPI/api/provider/yhdm"
)

type ApiProvider interface {
	//InitDefault()
	GetUniqueId() string
	Initialize() bool
	GetPlayUrls() []model.ApiResource
}

const (
	IMOMOE = "imomoe"
	YHDM   = "yhdm"
)

var ProviderStatusMap = map[string]bool{
	IMOMOE: true,
	YHDM:   true,
}

func IsProviderAvailable(provider string) bool {
	status, ok := ProviderStatusMap[provider]
	if ok {
		return status
	}
	return false
}

func InitWithUid(provider string, uid string) *ApiProvider {
	var pvdr ApiProvider
	switch provider {
	case IMOMOE:
		pvdr = imomoeApi.InitWithUid(uid)
	case YHDM:
		pvdr = yhdmApi.InitWithUid(uid)
	default:
		pvdr = imomoeApi.InitWithUid(uid)
	}
	return &pvdr
}
