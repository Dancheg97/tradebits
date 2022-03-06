package redis

import (
	"context"
	"time"
)

func Lock(key string) bool {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Millisecond,
	)
	defer cancel()
	blcmd := rds.SetNX(ctx, key, true, 0)
	wasSet, connErr := blcmd.Result()
	if connErr != nil {
		return false
	}
	return wasSet
}
