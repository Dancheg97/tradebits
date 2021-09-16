package lock

import (
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
