package auth_service

import "AynaAPI/server/models"

func GetPermission(username string) int {
	u, err := models.GetUser(username)
	if err != nil {
		return -1
	}
	if p, err := models.GetPermissionByUser(u); err == nil {
		return p.Level
	}
	return -1
}

func GetPermissionByUser(user *models.User) int {
	if p, err := models.GetPermissionByUser(user); err == nil {
		return p.Level
	}
	return -1
}
