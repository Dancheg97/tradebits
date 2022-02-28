package redis

import (
	"testing"
	"time"
)

func TestGenerateLockers(t *testing.T) {
	Setup("localhost:6397")
	resp := rds.SetNX(ctx, "x", "x", time.Millisecond)
	err := resp.Err()
	if err != nil {
		t.Error(err)
	}
}
