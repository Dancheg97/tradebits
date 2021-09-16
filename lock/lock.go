package lock

import (
	"errors"
	"sync"
	"time"
)

type blockedMap struct {
	mutex  sync.Mutex
	userId map[[64]byte]bool
}

func generateBlockers() map[byte]*blockedMap {
	var blockers = make(map[byte]*blockedMap)
	for i := 0; i < 256; i++ {
		var blockedMap = blockedMap{userId: make(map[[64]byte]bool)}
		blockers[byte(i)] = &blockedMap
	}
	return blockers
}

var blockers = generateBlockers()

func Lock(ID []byte) error {
	if len(ID) != 64 {
		return errors.New("error on checking length of bytes")
	}
	var lockID [64]byte
	copy(lockID[:], ID[:64])
	keyByte := ID[0]
	blocker := blockers[keyByte]
	blocker.mutex.Lock()
	_, found := blocker.userId[lockID]
	if found {
		blocker.mutex.Unlock()
		time.Sleep(time.Millisecond * 144)
		return Lock(ID)
	}
	blocker.userId[lockID] = true
	blocker.mutex.Unlock()
	return nil
}

func Unlock(ID []byte) {
	var lockID [64]byte
	copy(lockID[:], ID[:64])
	keyByte := ID[0]
	blocker := blockers[keyByte]
	blocker.mutex.Lock()
	defer blocker.mutex.Unlock()
	delete(blocker.userId, lockID)
}

//TODO implement recursive function lock when possible
