package api

import (
	"encoding/json"
	"net/http"
	"tradebits/mongo"
)

var MarketInfoResponse []byte

func InfoMarketGet(w http.ResponseWriter, r *http.Request) {
	w.Write(MarketInfoResponse)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
}

func InfoNetGet(w http.ResponseWriter, r *http.Request) {
	netmembers, errGetCollection := mongo.GetCollection("net")
	if errGetCollection != nil {
		w.WriteHeader(400)
		return
	}
	for _, member := range netmembers {
		delete(member, "_id")
	}
	respbytes, infoEjectErr := json.Marshal(netmembers)
	if infoEjectErr != nil {
		w.WriteHeader(400)
		return
	}
	w.Write(respbytes)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
}
