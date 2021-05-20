package main

import (
	"AynaAPI/config"
	"AynaAPI/server"
	"fmt"
)

func main() {
	router := server.InitRouter()
	err := router.Run(fmt.Sprintf(":%d", config.ServerConfig.HttpPort))
	if err != nil {
		return
	}
}
