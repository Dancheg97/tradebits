package redis

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestSetup(t *testing.T) {
	godotenv.Load("../.env")
	redis_host, _ := os.LookupEnv("redis_host")
	redis_name, _ := os.LookupEnv("redis_name")
	Setup(redis_host, redis_name)
	resp := rds.SetNX(context.Background(), "x", "x", time.Millisecond)
	err := resp.Err()
	if err != nil {
		t.Error(err)
	}
}
