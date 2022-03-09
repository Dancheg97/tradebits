package user

import (
	"encoding/json"
	"net/http"
)

func UserTradesGet(w http.ResponseWriter, r *http.Request) {
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
	tradeIds, err := mongo.FindIdx("trades", "ukey", ukey)
	if err != nil {
		w.WriteHeader(503)
		return
	}
	response := []map[string]interface{}{}
	for _, id := range tradeIds {
		trade := map[string]interface{}{}
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
