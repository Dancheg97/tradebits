package redis

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func getRedisHost() string {
	godotenv.Load("../.env")
	redis_host, _ := os.LookupEnv("REDIS_HOST")
	return redis_host
}

func TestSetup(t *testing.T) {
	err := Setup(getRedisHost())
	if err != nil {
		t.Error(err)
	}
}
