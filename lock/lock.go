package lock

func Lock(ID []byte) bool {
	blcmd := redisClient.SetNX(ctx, string(ID), true, 0)
	wasSet, connErr := blcmd.Result()
	if wasSet {
		return true
	}
	if connErr != nil {
		return false
	}
	return false
}
