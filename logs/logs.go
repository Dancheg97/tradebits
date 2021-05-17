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

func Response(w http.ResponseWriter, message string) {
	fmt.Printf("[GOOD-RESPONSE] %v\n", message)
	log.Printf("[GOOD-RESPONSE] %v\n", message)
	json.NewEncoder(w).Encode(message)
}

func ResponseErrString(w http.ResponseWriter, message string) {
	fmt.Printf("[ERROR-RESPONSE] %v\n", message)
	log.Printf("[ERROR-RESPONSE] %v\n", message)
	json.NewEncoder(w).Encode(message)
}

func Info(message string) {
	fmt.Printf("[INFO] %v\n", message)
	log.Printf("[INFO] %v\n", message)
}

func Warning(err error) {
	fmt.Printf("[WARNING] %v\n", err)
	log.Printf("[WARNING] %v\n", err)
}

func Error(err error) {
	fmt.Printf("[ERROR] %v\n", err)
	log.Printf("[ERROR] %v\n", err)
}

func Critical(err error) {
	fmt.Printf("[CRITICAL] %v\n", err)
	log.Printf("[CRITICAL] %v\n", err)
}
