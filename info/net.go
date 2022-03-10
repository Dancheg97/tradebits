package info

import "net/http"

func InfoMarketGet(w http.ResponseWriter, r *http.Request) {
	w.Write(info)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
}
