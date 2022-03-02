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
	redis_host, _ := os.LookupEnv("REDIS_HOST")
	redis_password, _ := os.LookupEnv("REDIS_PASSWORD")
	Setup(redis_host, redis_password)
	resp := rds.SetNX(context.Background(), "x", "x", time.Millisecond)
	err := resp.Err()
	if err != nil {
		t.Error(err)
	}
}
