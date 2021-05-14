package POST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type exchangerMessage struct {
	ExchangerPublicKey string `json:"ExchangerPublicKey"`
	ExchangerSign      string `json:"ExchangerSign"`
	OperationID        string `json:"OperationID"`
	ExchangerMessage   string `json:"ExchangerMessage"`
}

func ExchangerMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Message from exchanger")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var message exchangerMessage
	json.Unmarshal(requestBody, &message)
	var ExchangerPublicKey string = message.ExchangerPublicKey
	var ExchangerSign string = message.ExchangerSign
	var OperationID string = message.OperationID
	var exchangerMessage string = message.ExchangerMessage
	fmt.Println(ExchangerPublicKey, ExchangerSign, OperationID, exchangerMessage)
}

type userMessage struct {
	UserPublicKey string `json:"UserPublicKey"`
	UserSign      string `json:"UserSign"`
	OperationID   string `json:"OperationID"`
	UserMessage   string `json:"UserMessage"`
}

func UserMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Message from user")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var message userMessage
	json.Unmarshal(requestBody, &message)
	var UserPublicKey string = message.UserPublicKey
	var UserSign string = message.UserSign
	var OperationID string = message.OperationID
	var UserMessage string = message.UserMessage
	fmt.Println(UserPublicKey, UserSign, OperationID, UserMessage)
}
