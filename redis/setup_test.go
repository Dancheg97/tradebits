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
	Setup(redis_host)
	resp := rds.SetNX(context.Background(), "setuptest", "setuptest", time.Millisecond)
	err := resp.Err()
	if err != nil {
		t.Error(err)
	}
	Unlock("setuptest")
}
