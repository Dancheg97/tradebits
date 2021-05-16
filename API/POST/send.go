package POST

import (
	"bc_server/calc"
	"bc_server/database"
	"bc_server/lock"
	"encoding/base64"
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
		message := "json parse error"
		fmt.Println(message)
		fmt.Fprintf(w, message)
	}
	senderPublicKey := message.SenderPublicKey
	sendAmountBytes := message.SendAmountBytes
	recieverAdress := message.RecieverAdress
	senderSign := message.SenderSign
	// lock sender and reciever with defers to unlock
	senderAdress := calc.HashKey(senderPublicKey)
	lockErr := lockSenderAndReciever(senderAdress, recieverAdress)
	if lockErr != nil {
		message := "sender/reciever are locked with another transaction"
		fmt.Println(message)
		fmt.Fprintf(w, message)
		return
	}
	defer unlockSenderAndReciever(senderAdress, recieverAdress)
	// check balance
	sender, getSenderErr := database.GetUser(senderAdress)
	if getSenderErr != nil {
		message := "sender does not exist error"
		fmt.Println(message)
		fmt.Fprintf(w, message)
		return
	}
	sendAmount := binary.LittleEndian.Uint64(sendAmountBytes)
	balanceErr := checkBalance(&sender, sendAmount)
	if balanceErr != nil {
		message := "balance error"
		fmt.Println(message)
		fmt.Fprintf(w, message)
		return
	}
	// check sign
	messageArr := [][]byte{senderPublicKey, sendAmountBytes, recieverAdress}
	signErr := calc.Verify(messageArr, senderPublicKey, senderSign) //TODO ch
	if signErr != nil {
		message := "sign check error"
		fmt.Println(message)
		fmt.Fprintf(w, message)
		return
	}
	// TODO send transaction to syncronized nodes for verification, with a node sign
	// TODO allow all nodes to write transation to blockchain
	// TODO write transaction to blockchain
	// transfer money
	reciever, getRecieverErr := database.GetUser(recieverAdress)
	if getRecieverErr != nil {
		message := "get reciever error"
		fmt.Println(message)
		fmt.Fprintf(w, message)
		return
	}
	sender.SetMainBalance(sender.MainBalance - sendAmount)
	reciever.SetMainBalance(reciever.MainBalance + sendAmount)
	senderAdressBase64 := base64.RawStdEncoding.EncodeToString(senderAdress)
	recieverAdressBase64 := base64.RawStdEncoding.EncodeToString(recieverAdress)
	signBase64 := base64.RawStdEncoding.EncodeToString(senderSign)
	fmt.Printf("---\n[sender: %v]\n[send: %v]\n[reciever:%v]\n[sign:%v]\n---\n", senderAdressBase64, sendAmount, recieverAdressBase64, signBase64)
	fmt.Fprintf(w, "transaction passed")
}
