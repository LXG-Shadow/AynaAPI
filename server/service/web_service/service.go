package web_service

import (
	"AynaAPI/server/common"
	"AynaAPI/server/dao"
	"AynaAPI/server/service"
)

type WebService struct {
	service.Service
}

func New() *WebService {
	return &WebService{service.Service{
		Dao: dao.New(common.DBEngine),
	}}
}
