package gredis

import (
	"AynaAPI/config"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"time"
)

var Online bool = false
var RedisClient *redis.Client
var ctx = context.Background()

func Initialize() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Host,
		Password: config.RedisConfig.Password,
		DB:       0,
	})
	Online, _ = IsOnline()
}

func IsOnline() (bool, error) {
	if RedisClient == nil {
		return false, errors.New("redis client not initialize")
	}
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		return false, err
	}
	return true, nil
}

func IsExists(key string) bool {
	_, err := RedisClient.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		return false
	case err != nil:
		return false
	default:
		return true
	}
}

func Get(key string) (result *redis.StringCmd, ok bool) {
	result = RedisClient.Get(ctx, key)
	err := result.Err()
	switch {
	case err == redis.Nil:
		ok = false
	case err != nil:
		ok = false
	default:
		ok = true
	}
	return
}

func GetString(key string) (string, bool) {
	val, err := RedisClient.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		return val, false
	case err != nil:
		return val, false
	default:
		return val, true
	}
}

func GetData(key string, v interface{}) bool {
	if data, ok := GetString(key); ok {
		if err := json.Unmarshal([]byte(data), v); err != nil {
			return false
		}
		return true
	} else {
		return false
	}
}

func Set(key string, value interface{}, duration time.Duration) (bool, error) {
	if result := RedisClient.Set(ctx, key, value, duration); result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func SetData(key string, value interface{}, duration time.Duration) (bool, error) {
	strval, err := json.Marshal(value)
	if err != nil {
		return false, err
	}
	if result := RedisClient.Set(ctx, key, string(strval), duration); result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func Delete(key string) bool {
	if RedisClient.Del(ctx, key).Err() != nil {
		return false
	}
	return true
}
