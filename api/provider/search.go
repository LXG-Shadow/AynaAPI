package provider

import (
	"AynaAPI/api/model"
	imomoeApi "AynaAPI/api/provider/imomoe"
	susudmApi "AynaAPI/api/provider/susudm"
	yhdmApi "AynaAPI/api/provider/yhdm"
)

func Search(provider string, keyword string, page int) model.ApiResponse {
	switch provider {
	case IMOMOE:
		return imomoeApi.Search(keyword, page)
	case YHDM:
		return yhdmApi.Search(keyword, page)
	case SUSUDM:
		return susudmApi.Search(keyword, page)
	default:
		return susudmApi.Search(keyword, page)
	}
}
