package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"bc_server/API/GET"
	"bc_server/API/POST"
	"bc_server/calc"
	"bc_server/database"
)

/*
Подобно музыке поток идёт в глубины
Мгновенных операций стройный ряд
Мы любим жизнь, и любим горутины
Мы с ними строим райский сад
*/

func GetOwner() {
	//TODO replace with method that is connecting to other nodes to verify the increasing of balances
	file, _ := os.Open("public.pem")
	defer file.Close()
	keyData, _ := ioutil.ReadAll(file)
	keyStr := string(keyData)
	keyBytes, _ := base64.RawStdEncoding.DecodeString(keyStr)
	firstAdress := calc.HashKey(keyBytes)
	database.DB.Delete(firstAdress, nil)
	user, _ := database.NewUser(firstAdress)
	user.SetMainBalance(10000)
	adressBase64 := base64.RawStdEncoding.EncodeToString(firstAdress)
	fmt.Printf("---\n[start: :)]\n[reciever:%v]\n[amount:%v]\n---\n", adressBase64, 10000)
	time.Sleep(100000)
	recieverBytesAsStringBase64, _ := ioutil.ReadFile("tests/apiTests/public2.pem")
	recieverBytes, _ := base64.RawStdEncoding.DecodeString(string(recieverBytesAsStringBase64))
	recieverAdress := calc.HashKey(recieverBytes)
	database.DB.Delete(recieverAdress, nil)
	database.NewUser(recieverAdress)
}

func handleRequests() {
	http.HandleFunc("/send", POST.SendRequest)
	http.HandleFunc("/getbalance", GET.BalanceRequest)
	fmt.Printf("hostting on: localhost:8080\n")
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	GetOwner()
	handleRequests()
}
