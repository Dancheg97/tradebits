package logs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Init() {
	var file, _ = os.OpenFile("logs/info.log", os.O_CREATE|os.O_APPEND, 0644)
	log.SetOutput(file)
}

func Response(w http.ResponseWriter, message string) error {
	fmt.Printf("[GOOD-RESPONSE] %v\n", message)
	log.Printf("[GOOD-RESPONSE] %v\n", message)
	json.NewEncoder(w).Encode(message)
	return nil
}

func ResponseErrString(w http.ResponseWriter, message string) error {
	fmt.Printf("[ERROR-RESPONSE] %v\n", message)
	log.Printf("[ERROR-RESPONSE] %v\n", message)
	json.NewEncoder(w).Encode(message)
	return nil
}

func Info(message string) string {
	fmt.Printf("[INFO] %v\n", message)
	log.Printf("[INFO] %v\n", message)
	return message
}

func Warning(err error) error {
	fmt.Printf("[WARNING] %v\n", err)
	log.Printf("[WARNING] %v\n", err)
	return err
}

func Error(err error) error {
	fmt.Printf("[ERROR] %v\n", err)
	log.Printf("[ERROR] %v\n", err)
	return err
}

func Critical(err error) error {
	fmt.Printf("[CRITICAL] %v\n", err)
	log.Printf("[CRITICAL] %v\n", err)
	return err
}
