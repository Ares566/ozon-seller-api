package config

import (
	"github.com/go-redis/redis"
	"os"
)


// NewCacheConfig is
func NewCacheConfig() interface{} {

	if os.Getenv("CACHE_DRIVER") == "REDIS" {
		return &redis.Options{
			Addr:     os.Getenv("CACHE_HOST") + ":" + os.Getenv("CACHE_PORT"),
			Password: "",
			DB:       0,
		}
	}

	return nil
}
