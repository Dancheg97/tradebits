package user

import (
	"encoding/json"
	"net/http"
)

type Trade struct {
	Offer   int    `json:"offer"`
	Recieve int    `json:"recieve"`
	Mkey    string `json:"mkey"`
}

func UserTradesGet(w http.ResponseWriter, r *http.Request) {
	req := UkeyRequst{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(406)
		return
	}
	user := User{}
	notFound := mongo.Get("user", "ukey", req.Ukey, &user)
	if notFound != nil {
		w.WriteHeader(404)
		return
	}
	tradeIds, err := mongo.FindIdx("trades", "ukey", req.Ukey)
	if err != nil {
		w.WriteHeader(503)
		return
	}
	response := []Trade{}
	for _, id := range tradeIds {
		trade := Trade{}
		err := mongo.GetIdx("trades", id, &trade)
		if err != nil {
			w.WriteHeader(503)
			return
		}
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
