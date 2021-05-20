package main

import (
	"bc_server/calc"
	"bc_server/database"
	"bc_server/logs"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/send", POST.SendRequest)
	http.HandleFunc("/newuser", POST.NewUserRequest)
	http.HandleFunc("/getbalance", GET.BalanceRequest)
	fmt.Printf("hostting on: localhost:8080\n")
	http.ListenAndServe("localhost:8080", nil)
}

func Start() {
	pubBase64Bytes, _ := ioutil.ReadFile("public.pem")
	pubBaset64 := string(pubBase64Bytes)
	pubBytes, _ := base64.RawStdEncoding.DecodeString(pubBaset64)
	adress := calc.Hash(pubBytes)
	database.NewUser(adress)
	inputBalance := make([]byte, 8)
	binary.LittleEndian.PutUint64(inputBalance, 0)
	outputBalance := make([]byte, 8)
	binary.LittleEndian.PutUint64(outputBalance, 10000)
	input := calc.ConcatenateMessage([][]byte{adress, inputBalance})
	output := calc.ConcatenateMessage([][]byte{adress, outputBalance})
	transaction := database.Transaction{
		FirstObjectInput:  input,
		FirstObjectOutput: output,
		ShiftInfo:         outputBalance,
	}
	user, userGetErr := database.NewUser(adress)
	if userGetErr != nil {
		user, _ := database.GetUser(adress)
		user.SetMainBalance(10000)
	} else {
		user.SetMainBalance(10000)
	}
	transaction.WriteTransaction()
}

func main() {
	logs.Init()
	Start()
	handleRequests()
}
