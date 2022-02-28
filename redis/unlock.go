package redis

func Unlock(ID []byte) {
	rds.Del(ctx, string(ID))
}
