package POST

import (
	"bc_server/calc"
	"bc_server/database"
	"bc_server/logs"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type newUserRequest struct {
	SenderPublicKey []byte `json:"SenderPublicKey"`
	MessageKey      []byte `json:"MessageKey"`
	Image           []byte `json:"Image"`
	SenderSign      []byte `json:"SenderSign"`
}

func NewUserRequest(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message newUserRequest
	wrongRequest := json.Unmarshal(reqBody, &message)
	if wrongRequest != nil {
		logs.ResponseErrString(w, "json parse error")
		return
	}
	SenderPublicKey := message.SenderPublicKey
	MessageKey := message.MessageKey
	Image := message.Image
	SenderSign := message.SenderSign
	//check sign
	mes := [][]byte{SenderPublicKey, MessageKey, Image}
	signErr := calc.Verify(mes, SenderPublicKey, SenderSign)
	if signErr != nil {
		logs.ResponseErrString(w, "signature error")
		return
	}
	adress := calc.Hash(SenderPublicKey)
	user, userExistsErr := database.NewUser(adress)
	if userExistsErr != nil {
		logs.ResponseErrString(w, "user exists error")
		return
	}
	user.SetMessageKey(MessageKey)
	user.SetImage(Image)
	logs.Response(w, "user craeted")
}
