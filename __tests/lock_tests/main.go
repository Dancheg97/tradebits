package main

import (
	"sync_tree/__logs"
	"sync_tree/__tests"
	"sync_tree/_lock"
)

func lockUnlockTest() {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	err := _lock.Lock(lockBytes)
	if err != nil {
		__tests.Failed("lock", "lock", "attemt to lock and unlock user")
		return
	}
	_lock.Unlock(lockBytes)
	__tests.Passed("lock", "lock", "attemt to lock and unlock user")
}

func lockLockedTest() {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	_lock.Lock(lockBytes)
	err := _lock.Lock(lockBytes)
	if err != nil {
		__tests.Passed("lock", "lock", "attemt to lock locked user")
		return
	}
	__tests.Failed("lock", "lock", "attemt to lock locked user")
	_lock.Unlock(lockBytes)
}

func lockWrongIDTest() {
	lockBytes := make([]byte, 68)
	lockBytes[0] = 65
	lockBytes[1] = 66
	_lock.Lock(lockBytes)
	err := _lock.Lock(lockBytes)
	if err == nil {
		__tests.Failed("lock", "lock", "attemt to lock bad id")
		return
	}
	__tests.Passed("lock", "lock", "attemt to lock bad id")
}

func main() {
	__logs.Init()
	lockUnlockTest()
	lockLockedTest()
	lockWrongIDTest()
}
