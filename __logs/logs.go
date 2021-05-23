package __logs

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func Init() {
	var file, _ = os.OpenFile("__logs/info.log", os.O_CREATE|os.O_APPEND, 0644)
	log.SetOutput(file)
}

func Info(a ...interface{}) {
	log.Printf("[INFO] %v\n", fmt.Sprint(a))
}

func Warning(a ...interface{}) error {
	log.Printf("[WARN] %v\n", fmt.Sprint(a))
	return errors.New(fmt.Sprint(a))
}

func Error(a ...interface{}) error {
	log.Printf("[ERROR] %v\n", fmt.Sprint(a))
	return errors.New(fmt.Sprint(a))
}

func Critical(a ...interface{}) error {
	log.Printf("[FATAL] %v\n", fmt.Sprint(a))
	return errors.New(fmt.Sprint(a))
}
