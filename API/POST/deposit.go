package POST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type depositRequestStart struct {
	UserPublicKey   string `json:"UserPublicKey"`
	UserSign        string `json:"UserSign"`
	ExchangerAdress string `json:"ExchangerAdress"`
	OperationID     string `json:"OperationID"`
	DepositAmount   int64  `json:"DepositAmount"`
}

func DepositRequestStart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deposit approval")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var message depositRequestStart
	json.Unmarshal(requestBody, &message)
	var UserPublicKey string = message.UserPublicKey
	var UserSign string = message.UserSign
	var ExchangerAdress string = message.ExchangerAdress
	var OperationID string = message.OperationID
	var DepositAmount int64 = message.DepositAmount
	fmt.Println(UserPublicKey, UserSign, ExchangerAdress, OperationID, DepositAmount)
}

type depositRequestApproval struct {
	ExchangerPublicKey string `json:"ExchangerPublicKey"`
	ExchangerSign      string `json:"ExchangerSign"`
	OperationID        string `json:"OperationID"`
}

func DepositRequestApproval(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deposit request negation")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var message depositRequestApproval
	json.Unmarshal(requestBody, &message)
	var ExchangerPublicKey string = message.ExchangerPublicKey
	var ExchangerSign string = message.ExchangerSign
	var OperationID string = message.OperationID
	fmt.Println(ExchangerPublicKey, ExchangerSign, OperationID)
}

type depositRequestNegation struct {
	UserPublicKey string `json:"UserPublicKey"`
	UserSign      string `json:"UserSign"`
	OperationID   string `json:"OperationID"`
}

func DepositRequestNegation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deposit request bad end")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var message depositRequestNegation
	json.Unmarshal(requestBody, &message)
	var UserPublicKey string = message.UserPublicKey
	var UserSign string = message.UserSign
	var OperationID string = message.OperationID
	fmt.Println(UserPublicKey, UserSign, OperationID)
}
