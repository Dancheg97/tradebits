package POST

import (
	"bc_server/database"
	"bc_server/hash"
	"bc_server/lock"
	"bc_server/verify"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type sendRequest struct {
	SenderPublicKey []byte `json:"senderPublicKey"`
	Amount          []byte `json:"amount"`
	RecieverAdress  []byte `json:"user"`
	SenderSign      []byte `json:"sign"`
}

func lockParticipants(senderAdress []byte, recieverAdress []byte) {
	var senderAdressArray [64]byte
	var recieverAdressArray [64]byte
	copy(senderAdressArray[:], senderAdress[:64])
	copy(recieverAdressArray[:], recieverAdress[:64])
	lock.Lock(senderAdressArray)
	lock.Lock(recieverAdressArray)
}

func unlockParticipants(senderAdress []byte, recieverAdress []byte) {
	var senderAdressArray [64]byte
	var recieverAdressArray [64]byte
	copy(senderAdressArray[:], senderAdress[:64])
	copy(recieverAdressArray[:], recieverAdress[:64])
	lock.Unlock(senderAdressArray)
	lock.Unlock(recieverAdressArray)
}

func TransferMoney(senderAdress []byte, recieverAdress []byte, amount uint64) {
	senderBalance := database.ReadBalance(senderAdress)
	recieverBalance := database.ReadBalance(recieverAdress)
	database.WriteBalance(senderAdress, senderBalance-amount)
	database.WriteBalance(recieverAdress, recieverBalance-amount)
}

func CheckBalance(adress []byte, minimalBalance uint64) error {
	
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got send message")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message sendRequest
	json.Unmarshal(reqBody, &message)
	senderPublicKey := message.SenderPublicKey
	sendAmountBytes := message.Amount
	recieverAdress := message.RecieverAdress
	senderSign := message.SenderSign
	senderAdress := hash.HashKey(senderPublicKey)
	sendAmount := binary.LittleEndian.Uint64(sendAmountBytes)
	if senderBalance < sendAmount {
		json.NewEncoder(w).Encode("not enough balance")
		return
	}
	// step 2 check sign
	fullMessage := [][]byte{senderPublicKey, sendAmountBytes, recieverAdress}
	signError := verify.Verify(fullMessage, senderPublicKey, senderSign)
	if signError != nil {
		fmt.Println("sign is not ok")
		json.NewEncoder(w).Encode("sign verificaiton error")
		return
	}
	// step 3 - lock sender and reciever for processing time

	//step 4 - transfer money

	// step 5 - unlock sender and reciever

	json.NewEncoder(w).Encode("sucess")
}
