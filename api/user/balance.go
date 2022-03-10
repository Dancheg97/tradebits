package user

import (
	"encoding/json"
	"net/http"
)

type UkeyRequst struct {
	Ukey string `json:"ukey"`
}

type BalanceResponse struct {
	Balance int `json:"balance"`
}

func UserBalanceGet(w http.ResponseWriter, r *http.Request) {
	req := UkeyRequst{}
	json.NewDecoder(r.Body).Decode(&req)
	if req.Ukey == "" {
		w.WriteHeader(406)
		return
	}
	user := User{}
	notFound := mongo.Get("user", "ukey", req.Ukey, &user)
	if notFound != nil {
		w.WriteHeader(404)
		return
	}
	response := BalanceResponse{
		Balance: user.Balance,
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
