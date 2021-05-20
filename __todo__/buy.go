package POST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type buyRequest struct {
	UserPublicKey string `json:"UserPublicKey"`
	UserSign      string `json:"UserSign"`
	Asset         string `json:"Asset"`
	Amount        int64  `json:"Amount"`
	Direction     bool   `json:"Direction"`
}

func BuyMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got buy message")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message buyRequest
	json.Unmarshal(reqBody, &message)
	var UserPublicKey string = message.UserPublicKey
	var UserSign string = message.UserSign
	var Asset string = message.Asset
	var Amount int64 = message.Amount
	var Direction bool = message.Direction
	/*
		TODO - next sequence:
		check direction (true - daco to asset/ false - asset to daco)
		check balance
		check sign
	*/
	fmt.Println(UserPublicKey, UserSign, Asset, Amount, Direction)
}
