package GET

import (
	"bc_server/database"
	"bc_server/logs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type balanceRequest struct {
	Adress []byte `json:"adress"`
}

func BalanceRequest(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message balanceRequest
	wrongRequest := json.Unmarshal(reqBody, &message)
	if wrongRequest != nil {
		logs.ResponseErrString(w, "json parse error")
		return
	}
	user, getSenderErr := database.GetUser(message.Adress)
	if getSenderErr != nil {
		logs.ResponseErrString(w, "user does not exist error")
		return
	}
	logs.Response(w, fmt.Sprint(user.MainBalance))
}
