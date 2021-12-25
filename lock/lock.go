package lock

import "time"

func Lock(ID []byte) bool {
	strId := string(ID)
	blcmd := redisClient.SetNX(ctx, strId, true, time.Second)
	if blcmd.Err() == nil {
		return false
	}
	return true
}
