package logs

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var file, _ = os.OpenFile("logs/info.log", os.O_CREATE|os.O_APPEND, 0644)

func Init() {
	log.SetOutput(file)
}

func Info(a ...interface{}) {
	log.Printf("[INFO] %v\n", fmt.Sprint(a...))
}

func Warning(a ...interface{}) error {
	log.Printf("[WARN] %v\n", fmt.Sprint(a...))
	return errors.New(fmt.Sprint(a...))
}

func Error(a ...interface{}) error {
	log.Printf("[ERROR] %v\n", fmt.Sprint(a...))
	return errors.New(fmt.Sprint(a...))
}

func Critical(a ...interface{}) error {
	log.Printf("[FATAL] %v\n", fmt.Sprint(a...))
	return errors.New(fmt.Sprint(a...))
}