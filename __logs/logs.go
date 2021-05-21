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

func Info(message string, a ...interface{}) string {
	converted := fmt.Sprintf(message, a)
	log.Printf("[INFO] %v\n", converted)
	return converted
}

func Warning(message string, a ...interface{}) error {
	converted := fmt.Sprintf(message, a)
	log.Printf("[WARN] %v\n", converted)
	return errors.New(converted)
}

func Error(message string, a ...interface{}) error {
	converted := fmt.Sprintf(message, a)
	log.Printf("[ERROR] %v\n", converted)
	return errors.New(converted)
}

func Fatal(message string, a ...interface{}) error {
	converted := fmt.Sprintf(message, a)
	log.Printf("[FATAL] %v\n", converted)
	return errors.New(converted)
}
