package lock

import (
	"sync"
)

type blockedMap struct {
	mutex  sync.Mutex
	userId map[[64]byte]bool
}

func generateBlockers() map[byte]*blockedMap {
	var blockers = make(map[byte]*blockedMap)
	for i := 0; i < 256; i++ {
		var blockedMap = blockedMap{
			userId: make(map[[64]byte]bool),
		}
		blockers[byte(i)] = &blockedMap
	}
	return blockers
}

var blockers = generateBlockers()
