package user

import (
	"encoding/json"
	"net/http"
)

func UserBalanceGet(w http.ResponseWriter, r *http.Request) {
	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)
	ukey, exists := request["ukey"]
	if !exists {
		w.WriteHeader(406)
		return
	}
	user := User{}
	notFound := mongo.Get("user", "ukey", ukey, &user)
	if notFound != nil {
		w.WriteHeader(404)
		return
	}
	response := map[string]int{
		"balance": user.Balance,
	}
	respbytes, marshErr := json.Marshal(response)
	if marshErr != nil {
		w.WriteHeader(503)
		return
	}
	w.Write(respbytes)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
}
