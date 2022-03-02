package redis

import (
	"github.com/go-redis/redis/v8"
)

var rds *redis.Client

func Setup(host string, pass string) {
	rds = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass,
		DB:       0,
	})
}
