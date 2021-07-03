package models

import (
	"errors"
	"log"
)

type User struct {
	BaseModel
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}

type Permission struct {
	BaseModel
	UserID int
	Level  int
	User   User `gorm:"foreignKey:UserID;references:ID"`
}

func init() {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
		return
	}
	err = db.AutoMigrate(&Permission{})
	if err != nil {
		log.Println(err)
		return
	}
}

func IsUserNameExists(username string) (bool, error) {
	var user User
	err := db.Where(User{Username: username}).First(&user).Error
	if err != nil {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
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

func RegisterUser(username, password string) (*User, error) {
	if ok, _ := IsUserNameExists(username); ok {
		return nil, errors.New("username exists")
	}
	user := &User{Username: username, Password: password}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(username string) (*User, error) {
	var user User
	err := db.Where(User{Username: username}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func SetPermission(user *User) (*Permission, error) {
	perm := &Permission{UserID: user.ID, Level: 0}
	if err := db.Create(perm).Error; err != nil {
		return nil, err
	}
	return perm, nil
}

func GetPermissionByUser(user *User) (*Permission, error) {
	var permission Permission
	err := db.Where(Permission{UserID: user.ID}).First(&permission).Error
	if err != nil {
		return SetPermission(user)
	}
	return &permission, nil
}
