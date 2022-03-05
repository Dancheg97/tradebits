package redis

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func testSetup() {
	godotenv.Load("../.env")
	redis_host, _ := os.LookupEnv("REDIS_HOST")
	Setup(redis_host)
}

func TestSetup(t *testing.T) {
	testSetup()

	Unlock("setuptest")
}
