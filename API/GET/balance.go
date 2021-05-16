package GET

import (
	"bc_server/database"
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
		message := "json parse error"
		fmt.Println(message)
		fmt.Fprintf(w, message)
	}
	user, getSenderErr := database.GetUser(message.Adress)
	if getSenderErr != nil {
		message := "user does not exist error"
		fmt.Println(message)
		fmt.Fprintf(w, message)
		return
	}
	json.NewEncoder(w).Encode(user.MainBalance)
}
