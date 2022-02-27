package redis

import (
	"context"
	"testing"
	"time"
)

func TestGenerateLockers(t *testing.T) {
	Setup("localhost:6397")
	rds.SetNX(context.Background(), "x", "x", time.Millisecond)
}
