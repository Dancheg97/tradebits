package lock

import (
	"sync"
	"testing"
	"time"
)

func TimeOutError(t *testing.T) {
	time.Sleep(time.Second * 5)
	t.Error("timeout error for locking function")
}

func TestLockID(t *testing.T) {
	go TimeOutError(t)
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 69
	Lock(lockBytes)
	Unlock(lockBytes)
}

func HelperDefferedUnlock(IDtoUnlock []byte) {
	time.Sleep(time.Second)
	Unlock(IDtoUnlock)
}

func TestTryToLockLocked(t *testing.T) {
	go TimeOutError(t)
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	Lock(lockBytes)
	go HelperDefferedUnlock(lockBytes)
	Lock(lockBytes)
	defer Unlock(lockBytes)
}

type loka struct {
	mu   sync.Mutex
	cola bool
}

var operatedFine = 0

func LockAndUnlockWithTimer(lk *loka, t *testing.T) {
	lk.mu.Lock()
	time.Sleep(time.Second)
	lk.mu.Unlock()
	operatedFine += 1
}

func TestLockMultipleRequests(t *testing.T) {
	lk := loka{}
	go LockAndUnlockWithTimer(&lk, t)
	go LockAndUnlockWithTimer(&lk, t)
	go LockAndUnlockWithTimer(&lk, t)
	time.Sleep(time.Second * 5)
	if operatedFine != 3 {
		t.Error("some locking stuff")
	}
}
