package lock

func Unlock(ID []byte) {
	var lockID [64]byte
	copy(lockID[:], ID[:64])
	keyByte := ID[0]
	blocker := blockers[keyByte]
	blocker.mutex.Lock()
	defer blocker.mutex.Unlock()
	delete(blocker.userId, lockID)
}
