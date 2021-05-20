package POST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type assetMessage struct {
	AssetPublicKey string `json:"AssetPublicKey"`
	AssetSign      string `json:"AssetSign"`
	OperationID    string `json:"OperationID"`
	AssetMessage   string `json:"AssetMessage"`
}

func AssetMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Message from asset")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var message assetMessage
	json.Unmarshal(requestBody, &message)
	var AssetPublicKey string = message.AssetPublicKey
	var AssetSign string = message.AssetSign
	var OperationID string = message.OperationID
	var assetMessage string = message.AssetMessage
	fmt.Println(AssetPublicKey, AssetSign, OperationID, assetMessage)
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
