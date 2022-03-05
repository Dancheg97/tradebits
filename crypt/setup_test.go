package crypt

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestSetup(t *testing.T) {
	godotenv.Load("../.env")
	private, _ := os.LookupEnv("MARKET_PRIVATEKEY")
	err := Setup(private)
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

