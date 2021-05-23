package main

import (
	"fmt"
	"sync_tree/_lock"
)

func lockUblockTest() {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	err := _lock.Lock(lockBytes)
	if err != nil {
		fmt.Println("TEST NOT PASSED")
	}
	_lock.Unlock(lockBytes)
}

func main() {
	lockUblockTest()
	fmt.Println()
}
