package auth_service

import (
	"AynaAPI/server/model"
)

type LoginParam struct {
	Username string `form:"username" binding:"required,min=2,max=32"`
	Password string `form:"password" binding:"required,min=6,max=32"`
}

func (s *AuthService) AuthUser(param LoginParam) (*model.User, bool) {
	user, err := s.Dao.GetUser(param.Username)
	if err != nil {
		return nil, false
	}
	if user.Password != param.Password {
		return nil, false
	}
	return user, true
}

func (s *AuthService) Login(param LoginParam) bool {
	_, ok := s.AuthUser(param)
	return ok
}

func (s *AuthService) IsUserNameExists(username string) (bool, error) {
	user, err := s.Dao.GetUser(username)
	if err != nil {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}
