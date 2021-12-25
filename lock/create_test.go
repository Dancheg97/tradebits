package lock

import (
	"testing"
	"time"
)

func TestGenerateLockers(t *testing.T) {
	value := redisClient.SetNX(ctx, "ss", "ss", time.Millisecond)
	t.Error(value)
}
