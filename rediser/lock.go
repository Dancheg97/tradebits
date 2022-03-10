package rediser

import (
	"context"
	"errors"
	"time"
)

func (r *rediser) Lock(key string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Millisecond,
	)
	defer cancel()
	blcmd := r.db.SetNX(ctx, key, true, 0)
	success, connErr := blcmd.Result()
	if connErr == nil && success {
		return nil
	}
	return errors.New(connErr.Error() + "unable to lock: " + key)
}
