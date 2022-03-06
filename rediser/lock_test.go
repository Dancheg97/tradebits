package redis

import (
	"testing"
)

func TestLock(t *testing.T) {
	Setup(getRedisHost())
	rez := Lock("locktest1")
	if !rez {
		t.Error("value should be locked on first iteration")
	}
	Unlock("locktest1")
}

func TestLockLocked(t *testing.T) {
	Setup(getRedisHost())
	Lock("locktest2")
	rez := Lock("locktest2")
	if rez {
		t.Error("value should be unlocked on first iteration")
	}
	Unlock("locktest")
}
