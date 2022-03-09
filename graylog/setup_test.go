package graylog

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestSetup(t *testing.T) {
	godotenv.Load("../.env")
	grayApi, _ := os.LookupEnv("GRAYLOG_API")
	err := Setup(grayApi, 15)
	if err != nil {
		t.Error(err)
	}
}
