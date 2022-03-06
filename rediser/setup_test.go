package rediser

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func getRedis() (*rediser, error) {
	godotenv.Load("../.env")
	redis_host, _ := os.LookupEnv("REDIS_HOST")
	return Get(redis_host)
}

func TestSetup(t *testing.T) {
	_, err := getRedis()
	if err != nil {
		t.Error(err)
	}
}
