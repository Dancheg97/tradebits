package lock

import (
	"bc_server/logs"
	"errors"
	"fmt"
	"sync"
)

/*
переделать в блок по первому символу из кодировки base64
*/

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
		return logs.Error(errors.New("wrong bytes length"))
	}
	return nil
}

func Lock(ID []byte) error {
	fmt.Println("lock start")
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
		return errors.New("user already locked")
	}
	blocker.userId[lockID] = true
	return nil
}

func Unlock(ID []byte) {
	fmt.Println("lock end")
	var lockID [64]byte
	copy(lockID[:], ID[:64])
	keyByte := ID[0]
	blocker := blockers[keyByte]
	blocker.mutex.Lock()
	defer blocker.mutex.Unlock()
	delete(blocker.userId, lockID)
}
