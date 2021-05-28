package _lock

import (
	"testing"
	"sync_tree/__logs"
)

func TestLock(t *testing.T) {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	err := Lock(lockBytes)
	if err != nil {
		t.Error("lock", "lock", "attemt to lock and unlock user")
		__logs.TestFailed("lock", "lock", "attemt to lock and unlock user", t)
		return
	}
	Unlock(lockBytes)
	t.Error("lock", "lock", "attemt to lock and unlock user")
	__logs.TestPassed("lock", "lock", "attemt to lock and unlock user")
}

func TestUnlock(t *testing.T) {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	Lock(lockBytes)
	defer Unlock(lockBytes)
	err := Lock(lockBytes)
	if err != nil {
		__logs.TestPassed("lock", "lock", "attemt to lock locked user")
		return
	}
	__logs.TestFailed("lock", "lock", "attemt to lock locked user", t)
}

func TestOneMore(t *testing.T) {
	lockBytes := make([]byte, 68)
	lockBytes[0] = 65
	lockBytes[1] = 66
	Lock(lockBytes)
	err := Lock(lockBytes)
	if err == nil {
		__logs.TestFailed("lock", "lock", "attemt to lock locked user", t)
		return
	}
	__logs.TestPassed("lock", "lock", "attemt to lock bad id")
}
