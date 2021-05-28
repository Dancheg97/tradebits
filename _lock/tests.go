package sync_tree

import (
	"sync_tree/_lock"
	"testing"
)

func LockUnlockTest(t *testing.T) {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	err := _lock.Lock(lockBytes)
	if err != nil {
		TestFailed("lock", "lock", "attemt to lock and unlock user", t)
		return
	}
	_lock.Unlock(lockBytes)
	TestPassed("lock", "lock", "attemt to lock and unlock user")
}

func LockLockedTest(t *testing.T) {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	_lock.Lock(lockBytes)
	defer _lock.Unlock(lockBytes)
	err := _lock.Lock(lockBytes)
	if err != nil {
		TestPassed("lock", "lock", "attemt to lock locked user")
		return
	}
	TestFailed("lock", "lock", "attemt to lock locked user", t)
}

func LockWrongIDTest(t *testing.T) {
	lockBytes := make([]byte, 68)
	lockBytes[0] = 65
	lockBytes[1] = 66
	_lock.Lock(lockBytes)
	err := _lock.Lock(lockBytes)
	if err == nil {
		TestFailed("lock", "lock", "attemt to lock locked user", t)
		return
	}
	TestPassed("lock", "lock", "attemt to lock bad id")
}
