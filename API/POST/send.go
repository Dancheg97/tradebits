package POST

import (
	"bc_server/calc"
	"bc_server/database"
	"bc_server/lock"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type sendRequest struct {
	SenderPublicKey []byte `json:"senderPublicKey"`
	SendAmountBytes []byte `json:"amount"`
	RecieverAdress  []byte `json:"user"`
	SenderSign      []byte `json:"sign"`
	NodePublicKey   []byte `json:"NodePublicKey"`
	NodeSign        []byte `json:"NodeSign"`
}

func lockSenderAndReciever(sender []byte, reciever []byte) error {
	senderLockErr := lock.Lock(sender)
	if senderLockErr != nil {
		return errors.New("sender locked (for antoher operation)")
	}
	recieverLockErr := lock.Lock(reciever)
	if recieverLockErr != nil {
		lock.Unlock(sender)
		return errors.New("sender locked (for antoher operation)")
	}
	return nil
}

func unlockSenderAndReciever(sender []byte, reciever []byte) {
	lock.Unlock(sender)
	lock.Unlock(reciever)
}

func checkBalance(user *database.User, minValue uint64) error {
	if user.MainBalance < minValue {
		return errors.New("bad balance error")
	}
	return nil
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got send message")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message sendRequest
	json.Unmarshal(reqBody, &message)
	senderPublicKey := message.SenderPublicKey
	sendAmountBytes := message.SendAmountBytes
	recieverAdress := message.RecieverAdress
	senderSign := message.SenderSign
	// lock sender and reciever with defers to unlock
	senderAdress := calc.HashKey(senderPublicKey)
	lockErr := lockSenderAndReciever(senderAdress, recieverAdress)
	if lockErr != nil {
		json.NewEncoder(w).Encode(lockErr)
		return
	}
	defer unlockSenderAndReciever(senderAdress, recieverAdress)
	// check balance
	sender, getSenderErr := database.GetUser(senderAdress)
	if getSenderErr != nil {
		json.NewEncoder(w).Encode("sender does not exist")
		return
	}
	sendAmount := binary.LittleEndian.Uint64(sendAmountBytes)
	balanceErr := checkBalance(&sender, sendAmount)
	if balanceErr != nil {
		json.NewEncoder(w).Encode("not enough money to send")
		return
	}
	// check sign
	messageArr := [][]byte{senderPublicKey, sendAmountBytes, recieverAdress}
	signErr := calc.Verify(messageArr, senderAdress, senderSign)
	if signErr != nil {
		json.NewEncoder(w).Encode("bad sign")
		return
	}
	// TODO send transaction to syncronized nodes for verification, with a node sign
	// TODO allow all nodes to write transation to blockchain
	// TODO write transaction to blockchain
	// transfer money
	reciever, getRecieverErr := database.GetUser(recieverAdress)
	if getRecieverErr != nil {
		json.NewEncoder(w).Encode("reciever does not exist")
		return
	}
	sender.SetMainBalance(sender.MainBalance - sendAmount)
	reciever.SetMainBalance(reciever.MainBalance + sendAmount)
	json.NewEncoder(w).Encode("sucess")
}
