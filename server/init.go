package server

import (
	"AynaAPI/config"
	"AynaAPI/pkg/gredis"
	"AynaAPI/server/common"
	"AynaAPI/server/model"
	"github.com/gin-gonic/gin"
)

func Initialize() (err error) {
	err = common.Initialize()
	if err != nil {
		return
	}
	if config.ServerConfig.UseRedisCache {
		gredis.Initialize()
	}
	err = model.CreateTable()
	if err != nil {
		return
	}
	common.Logger.Infof("Redis Status: %t", gredis.Online)
	gin.SetMode(config.ServerConfig.GinMode)
	common.Logger.Info("Server Initialize success")
	return
}
