package config

import (
	"github.com/go-ini/ini"
	"log"
)

type API struct {
	Version           string
	Bilibili_SESSDATA string
	Bilibili_JCT      string
}

type Server struct {
	Version  string
	GinMode  string
	HttpPort int
}

var APIConfig *API
var ServerConfig *Server

func init() {
	APIConfig = &API{}
	ServerConfig = &Server{}
	Initialize()
}

func Initialize() {
	cfg, err := ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/conf.ini': %v", err)
	}

	mapTo(cfg, "API", APIConfig)
	mapTo(cfg, "Server", ServerConfig)
}

// mapTo map section
func mapTo(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
