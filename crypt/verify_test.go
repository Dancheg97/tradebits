package crypt

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestVerify(t *testing.T) {
	godotenv.Load("../.env")
	private, _ := os.LookupEnv("MARKET_PRIVATEKEY")
	crypter, _ := Crypter(private)
	success := crypter.Verify(testmes, crypter.pub, testsign)
	if !success {
		t.Error("sign was not verified")
	}
}
