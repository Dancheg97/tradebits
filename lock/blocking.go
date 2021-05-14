package lock

import (
	"sync"
	"time"
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

func Lock(ID [64]byte) {
	for {
		keyByte := ID[0]
		blocker := blockers[keyByte]
		blocker.mutex.Lock()
		if _, found := blocker.userId[ID]; !found {
			blocker.userId[ID] = true
			blocker.mutex.Unlock()
			return
		}
		time.Sleep(time.Millisecond * 8)
	}
}

func Unlock(ID [64]byte) {
	keyByte := ID[0]
	blocker := blockers[keyByte]
	blocker.mutex.Lock()
	delete(blocker.userId, ID)
	blocker.mutex.Unlock()
}
