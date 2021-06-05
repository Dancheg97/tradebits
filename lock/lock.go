package lock

import (
	"errors"
	"sync"
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

func checkLen(bytes []byte) error {
	if len(bytes) != 64 {
		return errors.New("error on checking length of bytes")
	}
	return nil
}

func Lock(ID []byte) error {
	lengthErr := checkLen(ID)
	if lengthErr != nil {
		return lengthErr
	}
	var lockID [64]byte
	copy(lockID[:], ID[:64])
	keyByte := ID[0]
	blocker := blockers[keyByte]
	blocker.mutex.Lock()
	defer blocker.mutex.Unlock()
	_, found := blocker.userId[lockID]
	if found {
		return errors.New("this key is already locked")
	}
	blocker.userId[lockID] = true
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
