package redis

func Unlock(ID []byte) {
	redisClient.Del(ctx, string(ID))
}
