package redis

import (
	"testing"
)

func TestUnlock(t *testing.T) {
	Setup(getRedisHost())
	Lock("unlocktest")
	rez := Unlock("unlocktest")
	if !rez {
		t.Error("value should be successfully unlcoked")
	}
}
