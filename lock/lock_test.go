package lock

import (
	"testing"
)

func TestLockID(t *testing.T) {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	err := Lock(lockBytes)
	if err != nil {
		t.Error("failed to lock and unlock id")
		return
	}
	Unlock(lockBytes)
}

func TestLockLockedID(t *testing.T) {
	lockBytes := make([]byte, 64)
	lockBytes[0] = 65
	lockBytes[1] = 66
	Lock(lockBytes)
	defer Unlock(lockBytes)
	err := Lock(lockBytes)
	if err != nil {
		return
	}
	t.Error("attemt to lock locked id succeded, test failed")
}

func TestWrongIDlength(t *testing.T) {
	lockBytes := make([]byte, 68)
	lockBytes[0] = 65
	lockBytes[1] = 66
	err := Lock(lockBytes)
	if err == nil {
		t.Error("attemt to lock bad id")
		return
	}
}
