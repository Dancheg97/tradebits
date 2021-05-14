package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"bc_server/distribution"
)

/*
Подобно музыке поток идёт в глубины
Мгновенных операций стройный ряд
Мы любим жизнь, и любим горутины
Мы с ними строим райский сад
*/

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/send", API.SendMessage)
	router.HandleFunc("/buy", API.BuyMessage)
	router.HandleFunc("/sell", API.SellMessage)
	router.HandleFunc("/deposit", API.DepositRequestStart)
	router.HandleFunc("/depositapproval", API.DepositRequestApproval)
	router.HandleFunc("/depositnegation", API.DepositRequestNegation)
	router.HandleFunc("/withdrawal", API.WithdrawalRequestStart)
	router.HandleFunc("/withdrawalapproval", API.WithdrawalRequestApproval)
	router.HandleFunc("/withdrawalnegation", API.WithdrawalRequestNegation)
	router.HandleFunc("/userMessage", API.UserMessage)
	router.HandleFunc("/exchangerMessage", API.ExchangerMessage)
	fmt.Printf("hostting on: localhost:8080\n")
	http.ListenAndServe("localhost:8080", router)
}

func main() {
	go distribution.StartDistribution()
	handleRequests()
}
