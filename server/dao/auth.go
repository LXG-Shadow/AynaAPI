package dao

import (
	"AynaAPI/server/model"
	"errors"
	"gorm.io/gorm"
)

func (d *Dao) GetUserById(userId int) (*model.User, error) {
	var user model.User
	err := d.engine.Where(model.User{BaseModel: model.BaseModel{ID: userId}}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *Dao) GetUser(username string) (*model.User, error) {
	var user model.User
	err := d.engine.Where(model.User{Username: username}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *Dao) CreateUser(username, password string) (*model.User, error) {
	if _, err := d.GetUser(username); err != nil {
		return nil, errors.New("username exists")
	}
	user := model.User{Username: username, Password: password}
	// create both permission and user, if any error occur, rollback.
	err := d.engine.Transaction(func(tx *gorm.DB) error {
		if err := d.engine.Create(&user).Error; err != nil {
			return err
		}
		permission := model.Permission{UserID: user.ID, Level: 0}
		if err := d.engine.Create(permission).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (d *Dao) GetPermission(userId int) (*model.Permission, error) {
	var permission model.Permission
	if err := d.engine.Where(model.Permission{UserID: userId}).First(&permission).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (d *Dao) SetPermission(userId, level int) (*model.Permission, error) {
	permission := model.Permission{UserID: userId}
	if err := d.engine.Model(&permission).Update("level", level).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}
