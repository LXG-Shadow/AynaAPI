package common

import (
	"AynaAPI/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func createDBEngine() (err error) {
	DBEngine, err = gorm.Open(sqlite.Open(config.ServerDBConfig.SqlitePath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.ServerDBConfig.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		Logger.WithField("error", err).Fatal("database connect fail")
		return
	}
	return
}
