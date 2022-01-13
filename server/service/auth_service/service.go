package auth_service

import (
	"AynaAPI/server/common"
	"AynaAPI/server/dao"
	"AynaAPI/server/service"
)

type AuthService struct {
	service.Service
}

func New() *AuthService {
	return &AuthService{service.Service{
		Dao: dao.New(common.DBEngine),
	}}
}
