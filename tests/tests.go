package main

import (
	"sync_tree/_lock"
	"testing"
)

func lockUnlockTest(t *testing.T) {
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

func lockLockedTest(t *testing.T) {
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

func lockWrongIDTest(t *testing.T) {
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
