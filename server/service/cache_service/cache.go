package cache_service

import (
	"AynaAPI/config"
	"time"
)

func GetCacheExpirePeriod() time.Duration {
	return time.Second * time.Duration(config.ServerConfig.RedisCachePeriod)
}
