package main

import (
	"AynaAPI/config"
	"AynaAPI/pkg/gredis"
	"AynaAPI/server"
	"fmt"
)

func main() {
	if config.ServerConfig.UseRedisCache {
		gredis.Initialize()
		fmt.Printf("redis status: %t\n", gredis.Online)
	}
	router := server.InitRouter()
	err := router.Run(fmt.Sprintf(":%d", config.ServerConfig.HttpPort))
	if err != nil {
		return
	}
}
