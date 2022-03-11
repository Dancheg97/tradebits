package graylog

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestSetInput(t *testing.T) {
	godotenv.Load("../.env")
	grayApi, _ := os.LookupEnv("GRAYLOG_API")
	err := setInput(grayApi)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckInput(t *testing.T) {
	godotenv.Load("../.env")
	grayApi, _ := os.LookupEnv("GRAYLOG_API")
	setInput(grayApi)
	err := checkInput(grayApi)
	if err != nil {
		t.Error(err)
	}
}
