package crypt

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func setupForTests() error {
	godotenv.Load("../.env")
	private, _ := os.LookupEnv("MARKET_PRIVATEKEY")
	return Setup(private)
}

func TestSetup(t *testing.T) {
	err := setupForTests()
	if err != nil {
		t.Error(err)
	}
	if priv == nil {
		t.Error("private key should not be nil")
	}
	if Pub == "" {
		t.Error("Public key should not be equal to default string")
	}
}
