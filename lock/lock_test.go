package lock

import (
	"testing"
	"time"
)

func TestLockID(t *testing.T) {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 69
	err := Lock(lockBytes)
	if err != nil {
		t.Error("failed to lock and unlock id")
		return
	}
	Unlock(lockBytes)
}

func HelperDefferedUnlock(IDtoUnlock []byte) {
	time.Sleep(time.Second)
	Unlock(IDtoUnlock)
}

func TestTryToLockLocked(t *testing.T) {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	Lock(lockBytes)
	go HelperDefferedUnlock(lockBytes)
	err := Lock(lockBytes)
	if err == nil {
		return
	}
	t.Error("attemt to lock locked id succeded, test failed")
}

func TestWrongIDlength(t *testing.T) {
	lockBytes := make([]byte, 68)
	lockBytes[0] = 65
	lockBytes[1] = 68
	err := Lock(lockBytes)
	if err == nil {
		t.Error("attemt to lock bad id")
		return
	}
}
