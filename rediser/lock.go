package rediser

import (
	"context"
	"time"
)

func (r *rediser) Lock(key string) bool {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Millisecond,
	)
	defer cancel()
	blcmd := r.db.SetNX(ctx, key, true, 0)
	wasSet, connErr := blcmd.Result()
	if connErr != nil {
		return false
	}
	return wasSet
}
