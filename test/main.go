package main

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var firstRedisClient = connectToRedis()
var secondResisClient = connectToRedis()

func connectToRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}

func tryToSetVal(inp string, ch chan<- bool) {
	rez := firstRedisClient.SetNX(ctx, inp, true, 0)
	written, _ := rez.Result()
	ch <- written
}

func main() {
	for i := 0; i < 1000000; i++ {
		trueflasechan := make(chan bool)
		someBytes := make([]byte, 6)
		rand.Reader.Read(someBytes)
		go tryToSetVal(string(someBytes), trueflasechan)
		go tryToSetVal(string(someBytes), trueflasechan)
		first := <-trueflasechan
		second := <-trueflasechan
		if first == second {
			fmt.Println("ERROR", first, second)
		}
	}
}
