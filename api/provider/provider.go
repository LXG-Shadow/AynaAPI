package provider

import (
	"AynaAPI/api/core"
	feijisuApi "AynaAPI/api/provider/feijisu"
	imomoeApi "AynaAPI/api/provider/imomoe"
	susudmApi "AynaAPI/api/provider/susudm"
	yhdmApi "AynaAPI/api/provider/yhdm"
	"encoding/json"
)

type ApiProvider interface {
	//InitDefault()
	GetUniqueId() string
	Initialize() bool
	GetPlayUrls() []core.ApiResource
}

const (
	IMOMOE  = "imomoe"
	YHDM    = "yhdm"
	SUSUDM  = "susudm"
	FEIJISU = "feijisu"
)

var ProviderStatusMap = map[string]bool{
	IMOMOE:  true,
	YHDM:    true,
	SUSUDM:  true,
	FEIJISU: true,
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
	case SUSUDM:
		pvdr, b = susudmApi.InitWithUrl(url)
	case FEIJISU:
		pvdr, b = feijisuApi.InitWithUrl(url)
	default:
		pvdr, b = susudmApi.InitWithUrl(url)
	}
	if b {
		return &pvdr, b
	}
	return nil, false
}

func InitWithJsonData(provider string, data []byte) (*ApiProvider, bool) {
	var pvdr ApiProvider
	var err error
	switch provider {
	case IMOMOE:
		var i imomoeApi.ImomoeVideo
		err = json.Unmarshal(data, &i)
		pvdr = &i
	case YHDM:
		var i yhdmApi.YhdmVideo
		err = json.Unmarshal(data, &i)
		pvdr = &i
	case SUSUDM:
		var i susudmApi.SusuDmVideo
		err = json.Unmarshal(data, &i)
		pvdr = &i
	case FEIJISU:
		var i feijisuApi.FeijisuDmVideo
		err = json.Unmarshal(data, &i)
		pvdr = &i
	default:
		var i susudmApi.SusuDmVideo
		err = json.Unmarshal(data, &i)
		pvdr = &i
	}
	if err == nil {
		return &pvdr, true
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
	case SUSUDM:
		pvdr, b = susudmApi.InitWithUid(uid)
	case FEIJISU:
		pvdr, b = feijisuApi.InitWithUid(uid)
	default:
		pvdr, b = susudmApi.InitWithUid(uid)
	}
	if b {
		return &pvdr, b
	}
	return nil, false
}
