package redis

import (
	"context"
	"time"
)

func Unlock(key string) bool {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Millisecond,
	)
	defer cancel()
	rez := rds.Del(ctx, string(key))
	return rez.Err() == nil
}
