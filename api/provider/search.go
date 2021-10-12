package provider

import (
	"AynaAPI/api/core"
	feijisuApi "AynaAPI/api/provider/feijisu"
	imomoeApi "AynaAPI/api/provider/imomoe"
	susudmApi "AynaAPI/api/provider/susudm"
	yhdmApi "AynaAPI/api/provider/yhdm"
)

func Search(provider string, keyword string, page int) core.ApiResponse {
	switch provider {
	case IMOMOE:
		return imomoeApi.Search(keyword, page)
	case YHDM:
		return yhdmApi.Search(keyword, page)
	case SUSUDM:
		return susudmApi.Search(keyword, page)
	case FEIJISU:
		return feijisuApi.Search(keyword, page)
	default:
		return susudmApi.Search(keyword, page)
	}
}
