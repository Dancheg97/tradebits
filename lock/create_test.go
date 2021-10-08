package lock

import (
	"sync"
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

func HelperDefferedUnlock(IDtoUnlock []byte) {
	time.Sleep(time.Second)
	Unlock(IDtoUnlock)
}

func TestTryToLockLocked(t *testing.T) {
	timeout := time.After(4 * time.Second)
	done := make(chan bool)
	go func() {
		lockBytes := make([]byte, 64)
		lockBytes[0] = 65
		lockBytes[1] = 66
		Lock(lockBytes)
		go HelperDefferedUnlock(lockBytes)
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

type loka struct {
	mu    sync.Mutex
	count int
}

func LockAndUnlockWithTimer(lk *loka, t *testing.T) {
	lk.mu.Lock()
	lk.count += 1
	time.Sleep(time.Second)
	lk.mu.Unlock()
}

func TestLockMultipleRequests(t *testing.T) {
	lk := loka{}
	go LockAndUnlockWithTimer(&lk, t)
	go LockAndUnlockWithTimer(&lk, t)
	go LockAndUnlockWithTimer(&lk, t)
	time.Sleep(time.Second * 5)
	if lk.count != 3 {
		t.Error("some locking stuff")
	}
}
