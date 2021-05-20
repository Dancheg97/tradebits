package POST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type sellRequest struct {
	UserPublicKey string `json:"UserPublicKey"`
	UserSign      string `json:"UserSign"`
	Asset         string `json:"Asset"`
	Amount        int64  `json:"Amount"`
	Price         int64  `json:"Price"`
	Direction     bool   `json:"Direction"`
}

func SellMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got sell message")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message sellRequest
	json.Unmarshal(reqBody, &message)
	var UserPublicKey string = message.UserPublicKey
	var UserSign string = message.UserSign
	var Asset string = message.Asset
	var Amount int64 = message.Amount
	var Price int64 = message.Price
	var Direction bool = message.Direction
	/*
		TODO - next sequence:
		check direction (true - daco to asset/ false - asset to daco)
		check balance
		check sign
		check orders on opposite side from low to high, if there are orders with price good enough price and volume, if yes close the offer, if no open new order, if yes block the orders id's and close by existings maby partially
	*/
	fmt.Println(UserPublicKey, UserSign, Asset, Amount, Price, Direction)
}
