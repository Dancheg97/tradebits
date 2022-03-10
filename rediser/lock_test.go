package rediser

import (
	"testing"
)

func TestLock(t *testing.T) {
	r, _ := getRedis()
	err := r.Lock("locktest1")
	if err != nil {
		t.Error(err)
	}
	r.Unlock("locktest1")
}

func TestLockLocked(t *testing.T) {
	r, _ := getRedis()
	r.Lock("locktest2")
	err := r.Lock("locktest2")
	if err == nil {
		t.Error(err)
	}
	r.Unlock("locktest")
}
