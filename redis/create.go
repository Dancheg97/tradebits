package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisClient = connectToRedis()

func connectToRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}