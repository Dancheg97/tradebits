package redis

import (
	"testing"
)

func TestLockID(t *testing.T) {
	rez1 := Lock("locktest")
	if !rez1 {
		t.Error("value should be locked on first iteration")
	}
	rez2 := Lock("locktest")
	if rez2 {
		t.Error("value should be unlocked on first iteration")
	}
	Unlock("locktest")
}
