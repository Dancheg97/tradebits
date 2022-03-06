package rediser

import (
	"testing"
)

func TestUnlock(t *testing.T) {
	r, _ := getRedis()
	r.Lock("unlocktest")
	rez := r.Unlock("unlocktest")
	if !rez {
		t.Error("value should be successfully unlcoked")
	}
}
