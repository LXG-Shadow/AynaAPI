package config

import (
	"github.com/go-ini/ini"
	"log"
	"path/filepath"
	"time"
)

type API struct {
	Version           string
	Bilibili_SESSDATA string
	Bilibili_JCT      string

	NovelRulePath string
}

type Server struct {
	Version   string
	GinMode   string
	HttpPort  int
	JwtSecret string
	FileRoot  string
	RealUrl   string

	UploadImageMaxSize int64

	UseRedisCache bool
}

func (self *Server) GetFilePath(path string) string {
	return filepath.Join(self.FileRoot, path)
}

type ServerDB struct {
	SqlitePath  string
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var APIConfig *API
var ServerConfig *Server
var ServerDBConfig *ServerDB
var RedisConfig *Redis

func init() {
	APIConfig = &API{}
	ServerConfig = &Server{}
	ServerDBConfig = &ServerDB{}
	RedisConfig = &Redis{}
	Initialize()
}

func Initialize() {
	cfg, err := ini.Load("conf/conf.ini")
	if err != nil {
		cfg, err = ini.Load("D:\\Repository\\AynaAPI\\conf\\test_conf.ini")
		log.Fatalf("setting.Setup, fail to parse 'conf/conf.ini': %v", err)
	}

	mapTo(cfg, "API", APIConfig)
	mapTo(cfg, "Server", ServerConfig)
	mapTo(cfg, "ServerDB", ServerDBConfig)
	mapTo(cfg, "Redis", RedisConfig)
}

// mapTo map section
func mapTo(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
