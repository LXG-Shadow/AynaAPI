package provider

import (
	"AynaAPI/api/model"
	imomoeApi "AynaAPI/api/provider/imomoe"
	yhdmApi "AynaAPI/api/provider/yhdm"
)

func Search(provider string, keyword string, page int) model.ApiResponse {
	switch provider {
	case IMOMOE:
		return imomoeApi.Search(keyword, page)
	case YHDM:
		return yhdmApi.Search(keyword, page)
	default:
		return imomoeApi.Search(keyword, page)
	}
}
