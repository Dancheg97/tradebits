package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rds *redis.Client

func Setup(adress string) {
	rds = redis.NewClient(&redis.Options{
		Addr:     adress,
		Password: "",
		DB:       0,
	})
}
