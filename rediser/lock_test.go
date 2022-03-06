package rediser

import (
	"testing"
)

func TestLock(t *testing.T) {
	r, _ := getRedis()
	rez := r.Lock("locktest1")
	if !rez {
		t.Error("value should be locked on first iteration")
	}
	r.Unlock("locktest1")
}

func TestLockLocked(t *testing.T) {
	r, _ := getRedis()
	r.Lock("locktest2")
	rez := r.Lock("locktest2")
	if rez {
		t.Error("value should be unlocked on first iteration")
	}
	r.Unlock("locktest")
}
