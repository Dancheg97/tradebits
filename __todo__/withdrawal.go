package POST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type withdrawalRequestStart struct {
	UserPublicKey string `json:"UserPublicKey"`
	UserSign      string `json:"UserSign"`
	AssetAdress   string `json:"AssetAdress"`
	OperationID   string `json:"OperationID"`
	Amount        int64  `json:"Amount"`
	Message       string `json:"Message"`
}

func WithdrawalRequestStart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got withdrawal request")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message withdrawalRequestStart
	json.Unmarshal(reqBody, &message)
	var UserPublicKey string = message.UserPublicKey
	var UserSign string = message.UserSign
	var AssetAdress string = message.AssetAdress
	var OperationID string = message.OperationID
	var Amount int64 = message.Amount
	var Message string = message.Message
	fmt.Println(UserPublicKey, UserSign, AssetAdress, OperationID, Amount, Message)
}

type withdrawalRequestApproval struct {
	UserPublicKey string `json:"UserPublicKey"`
	UserSign      string `json:"UserSign"`
	OperationID   string `json:"OperationID"`
}

func WithdrawalRequestApproval(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got withdrawal request")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message withdrawalRequestApproval
	json.Unmarshal(reqBody, &message)
	var UserPublicKey string = message.UserPublicKey
	var UserSign string = message.UserSign
	var OperationID string = message.OperationID
	fmt.Println(UserPublicKey, UserSign, OperationID)
}

type withdrawalRequestNegation struct {
	UserPublicKey string `json:"UserPublicKey"`
	UserSign      string `json:"UserSign"`
	OperationID   string `json:"OperationID"`
}

func WithdrawalRequestNegation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got withdrawal request")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message withdrawalRequestNegation
	json.Unmarshal(reqBody, &message)
	var UserPublicKey string = message.UserPublicKey
	var UserSign string = message.UserSign
	var OperationID string = message.OperationID
	fmt.Println(UserPublicKey, UserSign, OperationID)
}
