package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var rds *redis.Client

func Setup(host string) error {
	rds = redis.NewClient(&redis.Options{
		Addr: host,
		DB:   0,
	})
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	rds.SetNX(ctx, "setup", struct{}{}, time.Millisecond)
	resp := rds.Del(ctx, "setup")
	return resp.Err()
}
