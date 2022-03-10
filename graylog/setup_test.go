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

func TestCheckInput(t *testing.T) {
	godotenv.Load("../.env")
	grayApi, _ := os.LookupEnv("GRAYLOG_API")
	Setup(grayApi, 10)
	found, err := CheckInput(grayApi)
	if err != nil {
		t.Error(err)
	}
	if !found {
		t.Error("input not found")
	}
}
