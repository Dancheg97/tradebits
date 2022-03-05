package crypt

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestSetup(t *testing.T) {
	godotenv.Load("../.env")
	private, _ := os.LookupEnv("MARKET_PRIVATEKEY")
}
