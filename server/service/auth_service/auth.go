package auth_service

import "AynaAPI/server/models"

func Auth(username, password string) bool {
	if ok, err := models.AuthUser(username, password); err == nil {
		return ok
	}
	return false
}

func GetAuthUser(username, password string) (bool, *models.User) {
	user, err := models.GetUser(username)
	if err != nil {
		return false, nil
	}
	if user.Password != password {
		return false, nil
	}
	return true, user
}
