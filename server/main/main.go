package main

import (
	"AynaAPI/config"
	"AynaAPI/server"
	"flag"
	"fmt"
)

var configPath = flag.String("c", "conf/conf.ini", "-c config path")

func main() {
	flag.Parse()
	config.Load(*configPath)
	server.Initialize()
	router := server.InitRouter()
	err := router.Run(fmt.Sprintf(":%d", config.ServerConfig.HttpPort))
	if err != nil {
		return
	}
}
