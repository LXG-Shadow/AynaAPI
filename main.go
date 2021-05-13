package main

import (
	"AynaAPI/config"
	"AynaAPI/server"
	"fmt"
)

func main() {
	////fmt.Println(imomoe.Search("刀剑",1))
	//imomoe.GetInfo("855","0","0")
	//var v provider.ApiProvider= &imomoe.ImomoeVideo{
	//	Id:       "7599",
	//	SourceId: "0",
	//	EpId:     "0",
	//}
	//fmt.Println(v.Initialize())
	//
	//fmt.Println(v.GetPlayUrls())

	router := server.InitRouter()
	err := router.Run(fmt.Sprintf(":%d", config.ServerConfig.HttpPort))
	if err != nil {
		return
	}
}
