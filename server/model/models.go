package model

import (
	"AynaAPI/server/common"
	"github.com/sirupsen/logrus"
)

type BaseModel struct {
	ID int `gorm:"primaryKey" json:"id"`
}

func CreateTable() error {
	if err := createTable(&User{}, &Permission{}); err != nil {
		return err
	}
	if err := createTable(&UploadFile{}); err != nil {
		return err
	}
	return nil
}

func createTable(tables ...interface{}) error {
	for _, table := range tables {
		err := common.DBEngine.AutoMigrate(table)
		if err != nil {
			common.Logger.WithFields(logrus.Fields{
				"error": err,
			}).Warn("create table fail")
			return err
		}
	}
	return nil
}
