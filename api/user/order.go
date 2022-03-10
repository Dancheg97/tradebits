package user

import "net/http"

type Order struct {
	Ukey    string `bson:"ukey"`
	Mkey    string `bson:"mkey"`
	Offer   int    `bson:"offer"`
	Recieve int    `bson:"recieve"`
}

type OrderRequest struct {
	Hkey    string `json:"hkey"`
	Ukey    string `json:"ukey"`
	Mkey    string `json:"mkey"`
	Offer   int    `json:"offer"`
	Recieve int    `json:"recieve"`
	Sign    string `json:"sign"`
}

func UserOrderPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
