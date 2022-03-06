package rediser

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type IRediser interface {
	Lock(key string) bool
	Unlock(key string) bool
}

type rediser struct {
	db *redis.Client
}

func Get(host string) (*rediser, error) {
	rds := redis.NewClient(&redis.Options{
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
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	return &rediser{db: rds}, nil
}
