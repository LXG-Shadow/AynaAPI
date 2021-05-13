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

func InitWithUrl(provider string, url string) (*ApiProvider, bool) {
	var pvdr ApiProvider
	var b bool
	switch provider {
	case IMOMOE:
		pvdr, b = imomoeApi.InitWithUrl(url)
	case YHDM:
		pvdr, b = yhdmApi.InitWithUrl(url)
	default:
		pvdr, b = imomoeApi.InitWithUrl(url)
	}
	if b {
		return &pvdr, b
	}
	return nil, false
}

func InitWithUid(provider string, uid string) (*ApiProvider, bool) {
	var pvdr ApiProvider
	var b bool
	switch provider {
	case IMOMOE:
		pvdr, b = imomoeApi.InitWithUid(uid)
	case YHDM:
		pvdr, b = yhdmApi.InitWithUid(uid)
	default:
		pvdr, b = imomoeApi.InitWithUid(uid)
	}
	if b {
		return &pvdr, b
	}
	return nil, false
}
