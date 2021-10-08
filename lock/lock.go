package lock

import (
	"time"
)

func Lock(ID []byte) {
	var lockID [64]byte
	copy(lockID[:], ID[:64])
	keyByte := ID[0]
	blocker := blockers[keyByte]
	blocker.mutex.Lock()
	_, found := blocker.userId[lockID]
	if found {
		blocker.mutex.Unlock()
		time.Sleep(time.Millisecond * 144)
		Lock(ID)
		return
	}
	blocker.userId[lockID] = true
	blocker.mutex.Unlock()
}
