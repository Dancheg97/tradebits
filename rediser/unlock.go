package rediser

import (
	"context"
	"time"
)

func (r *rediser)Unlock(key string) bool {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Millisecond,
	)
	defer cancel()
	rez := r.db.Del(ctx, string(key))
	return rez.Err() == nil
}
