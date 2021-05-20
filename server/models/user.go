package models

import "log"

type User struct {
	BaseModel
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}

func init() {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
		return
	}
}

func AuthUser(username, password string) (bool, error) {
	var user User
	err := db.Where(User{Username: username}).First(&user).Error
	if err != nil {
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}
