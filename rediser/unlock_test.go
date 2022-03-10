package rediser

import (
	"testing"
)

func TestUnlock(t *testing.T) {
	r, _ := getRedis()
	r.Lock("unlocktest")
	err := r.Unlock("unlocktest")
	if err != nil {
		t.Error("value should be successfully unlcoked")
	}
}
