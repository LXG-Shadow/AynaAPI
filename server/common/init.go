package common

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Logger   *logrus.Logger
	DBEngine *gorm.DB
)

func Initialize() (err error) {
	err = createLogger()
	if err != nil {
		return
	}
	err = createDBEngine()
	if err != nil {
		return
	}
	return
}
