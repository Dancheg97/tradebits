package crypter

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestVerify(t *testing.T) {
	godotenv.Load("../.env")
	private, _ := os.LookupEnv("MARKET_PRIVATEKEY")
	crypter, _ := Get(private)
	err := crypter.Verify(testmes, crypter.pub, testsign)
	if err != nil {
		t.Error("sign was not verified")
	}
}
