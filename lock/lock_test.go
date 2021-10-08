package lock

import (
	"testing"
	"time"
)

func TestLockID(t *testing.T) {
	timeout := time.After(3 * time.Second)
	done := make(chan bool)
	go func() {
		lockBytes := make([]byte, 64)
		lockBytes[0] = 65
		lockBytes[1] = 69
		Lock(lockBytes)
		Unlock(lockBytes)
		done <- true
	}()
	select {
	case <-timeout:
		t.Fatal("Test didn't finish in time")
	case <-done:
	}
}

func TestTryToLockLocked(t *testing.T) {
	timeout := time.After(4 * time.Second)
	done := make(chan bool)
	go func() {
		lockBytes := make([]byte, 64)
		lockBytes[0] = 65
		lockBytes[1] = 66
		Lock(lockBytes)
		go func() {
			time.Sleep(time.Second)
			Unlock(lockBytes)
		}()
		Lock(lockBytes)
		Unlock(lockBytes)
		done <- true
	}()
	select {
	case <-timeout:
		t.Fatal("Test didn't finish in time")
	case <-done:
	}
}
