package user

import "net/http"

type Order struct {
	Ukey    string `bson:"ukey"`
	Mkey    string `bson:"mkey"`
	Offer   int    `bson:"offer"`
	Recieve int    `bson:"recieve"`
}

func UserOrderPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
