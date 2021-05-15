package main

import (
	"fmt"
	"net/http"

	"bc_server/API/POST"
	"bc_server/distribution"

	"github.com/gorilla/mux"
)

/*
Подобно музыке поток идёт в глубины
Мгновенных операций стройный ряд
Мы любим жизнь, и любим горутины
Мы с ними строим райский сад
*/

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/send", POST.SendMessage)
	router.HandleFunc("/buy", POST.BuyMessage)
	router.HandleFunc("/sell", POST.SellMessage)
	router.HandleFunc("/deposit", POST.DepositRequestStart)
	router.HandleFunc("/depositapproval", POST.DepositRequestApproval)
	router.HandleFunc("/depositnegation", POST.DepositRequestNegation)
	router.HandleFunc("/withdrawal", POST.WithdrawalRequestStart)
	router.HandleFunc("/withdrawalapproval", POST.WithdrawalRequestApproval)
	router.HandleFunc("/withdrawalnegation", POST.WithdrawalRequestNegation)
	router.HandleFunc("/userMessage", POST.UserMessage)
	router.HandleFunc("/exchangerMessage", POST.ExchangerMessage)
	fmt.Printf("hostting on: localhost:8080\n")
	http.ListenAndServe("localhost:8080", router)
}

func main() {
	go distribution.StartDistribution()
	handleRequests()
}
