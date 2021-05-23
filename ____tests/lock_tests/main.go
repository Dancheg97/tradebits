package main

import (
	"fmt"
	"sync_tree/__logs"
	"sync_tree/_lock"
)

func lockUnlockTest() {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	err := _lock.Lock(lockBytes)
	if err != nil {
		fmt.Println("\033[31m[TEST] (LOCK) {Lock} - failed\033[0m")
		return
	}
	_lock.Unlock(lockBytes)
	fmt.Println("\033[32m[TEST] (LOCK) {Lock} - passed\033[0m")
}

func lockLockedTest() {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	_lock.Lock(lockBytes)
	err := _lock.Lock(lockBytes)
	if err != nil {
		fmt.Println("\033[32m[TEST] (LOCK) {Lock locked} - passed\033[0m")
		return
	}
	fmt.Println("\033[31m[TEST] (LOCK) {Lock locked} - failed\033[0m")
	_lock.Unlock(lockBytes)
}

func lockWrongIDTest() {
	lockBytes := make([]byte, 68)
	lockBytes[0] = 65
	lockBytes[1] = 66
	_lock.Lock(lockBytes)
	err := _lock.Lock(lockBytes)
	if err == nil {
		fmt.Println("\033[31m[TEST] (LOCK) {Lock bad id} - failed\033[0m")
		return
	}
	fmt.Println("\033[32m[TEST] (LOCK) {Lock bad id} - passed\033[0m")
}

func main() {
	__logs.Init()
	lockUnlockTest()
	lockLockedTest()
	lockWrongIDTest()
}
