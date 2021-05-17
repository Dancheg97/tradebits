package POST

import (
	"bc_server/calc"
	"bc_server/database"
	"bc_server/lock"
	"bc_server/logs"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type sendRequest struct {
	SenderPublicKey []byte `json:"SenderPublicKey"`
	SendAmountBytes []byte `json:"SendAmountBytes"`
	RecieverAdress  []byte `json:"RecieverAdress"`
	SenderSign      []byte `json:"SenderSign"`
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

func SendRequest(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message sendRequest
	wrongRequest := json.Unmarshal(reqBody, &message)
	if wrongRequest != nil {
		logs.ResponseErrString(w, "json parse error")
		return
	}
	parseConditions := []bool{
		message.SenderPublicKey != nil,
		message.SendAmountBytes
	}
	senderPublicKey := message.SenderPublicKey
	sendAmountBytes := message.SendAmountBytes
	recieverAdress := message.RecieverAdress
	senderSign := message.SenderSign
	fmt.Println(senderPublicKey)
	fmt.Println(sendAmountBytes)
	fmt.Println(recieverAdress)
	fmt.Println(senderSign)
	// lock sender and reciever with defers to unlock
	senderAdress := calc.Hash(senderPublicKey)
	lockErr := lockSenderAndReciever(senderAdress, recieverAdress)
	if lockErr != nil {
		logs.ResponseErrString(w, "users are lcoked")
		return 
	}
	defer unlockSenderAndReciever(senderAdress, recieverAdress)
	// check balance
	sender, getSenderErr := database.GetUser(senderAdress)
	if getSenderErr != nil {
		logs.ResponseErrString(w, "sender does not exist error")
		return
	}
	sendAmount := binary.LittleEndian.Uint64(sendAmountBytes)
	balanceErr := checkBalance(&sender, sendAmount)
	if balanceErr != nil {
		logs.ResponseErrString(w, "balance error")
		return
	}
	// check sign
	messageArr := [][]byte{senderPublicKey, sendAmountBytes, recieverAdress}
	signErr := calc.Verify(messageArr, senderPublicKey, senderSign) //TODO ch
	if signErr != nil {
		logs.ResponseErrString(w, "sign check error")
		return
	}
	// TODO send transaction to syncronized nodes for verification, with a node sign
	// TODO allow all nodes to write transation to blockchain
	// TODO write transaction to blockchain
	// transfer money
	reciever, getRecieverErr := database.GetUser(recieverAdress)
	if getRecieverErr != nil {
		logs.ResponseErrString(w, "get reciever error")
		return
	}
	sender.SetMainBalance(sender.MainBalance - sendAmount)
	reciever.SetMainBalance(reciever.MainBalance + sendAmount)
	logs.Response(w, "transaction passed")
}
