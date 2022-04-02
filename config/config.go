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

	AnimeAgefansBaseUrl string
	AnimeOmofunBaseUrl  string
	AnimeDldmBaseUrl    string
}

type Server struct {
	Version      string
	GinMode      string
	HttpPort     int
	JwtSecret    string
	JwtTokenName string
	FileRoot     string
	RealUrl      string

	UploadSavePath       string
	UploadServerUrl      string
	UploadMaxSize        int
	UploadAllowImageExts []string

	UseRedisCache    bool
	RedisCachePeriod int // in seconds

	LogFile string
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
var cfgFile *ini.File

func init() {
	Load("conf/conf.ini")
}

func Load(path string) {
	var err error
	cfgFile, err = ini.Load(path)
	if err != nil {
		//log.Fatal("Load config fail")
		log.Println("Load config fail")
		return
	}
	Initialize()
}

func Initialize() {
	APIConfig = &API{}
	ServerConfig = &Server{}
	ServerDBConfig = &ServerDB{}
	RedisConfig = &Redis{}
	mapTo(cfgFile, "API", APIConfig)
	mapTo(cfgFile, "Server", ServerConfig)
	mapTo(cfgFile, "ServerDB", ServerDBConfig)
	mapTo(cfgFile, "Redis", RedisConfig)
}

// mapTo map section
func mapTo(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
