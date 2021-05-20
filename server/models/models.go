package models

import (
	"AynaAPI/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

type BaseModel struct {
	ID int `gorm:"primaryKey" json:"id"`
}

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open(config.ServerDBConfig.SqlitePath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.ServerDBConfig.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println(err)
	}
}
