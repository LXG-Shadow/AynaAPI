package auth_service

import "AynaAPI/server/model"

func (s *AuthService) GetPermissionByUser(user *model.User) int {
	if p, err := s.Dao.GetPermission(user.ID); err == nil {
		return p.Level
	}
	return -1
}
