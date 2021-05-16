package main

import (
	"fmt"
	"net/http"

	"bc_server/API/GET"
	"bc_server/API/POST"
)


func handleRequests() {
	http.HandleFunc("/send", POST.SendRequest)
	http.HandleFunc("/getbalance", GET.BalanceRequest)
	fmt.Printf("hostting on: localhost:8080\n")
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	handleRequests()
}
