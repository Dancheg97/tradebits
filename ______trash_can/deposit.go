package POST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type depositRequestStart struct {
	UserPublicKey string `json:"UserPublicKey"`
	UserSign      string `json:"UserSign"`
	AssetAdress   string `json:"AssetAdress"`
	OperationID   string `json:"OperationID"`
	DepositAmount int64  `json:"DepositAmount"`
}

func DepositRequestStart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deposit approval")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var message depositRequestStart
	json.Unmarshal(requestBody, &message)
	var UserPublicKey string = message.UserPublicKey
	var UserSign string = message.UserSign
	var AssetAdress string = message.AssetAdress
	var OperationID string = message.OperationID
	var DepositAmount int64 = message.DepositAmount
	fmt.Println(UserPublicKey, UserSign, AssetAdress, OperationID, DepositAmount)
}

type depositRequestApproval struct {
	AssetPublicKey string `json:"AssetPublicKey"`
	AssetSign      string `json:"AssetSign"`
	OperationID    string `json:"OperationID"`
}

func DepositRequestApproval(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deposit request negation")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var message depositRequestApproval
	json.Unmarshal(requestBody, &message)
	var AssetPublicKey string = message.AssetPublicKey
	var AssetSign string = message.AssetSign
	var OperationID string = message.OperationID
	fmt.Println(AssetPublicKey, AssetSign, OperationID)
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
