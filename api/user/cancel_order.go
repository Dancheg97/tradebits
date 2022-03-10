package user

import (
	"encoding/json"
	"net/http"
)

type CancelOrderRequest struct {
	Hkey string `json:"hkey"`
	Ukey string `json:"ukey"`
	Mkey string `json:"mkey"`
	Sign string `json:"sign"`
}

func UserCancelordersPost(w http.ResponseWriter, r *http.Request) {
	req := CancelOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(406)
		return
	}
	if req.Hkey != crypt.Pub() {
		w.WriteHeader(421)
		return
	}
	verfied := crypt.Verify(req.Hkey+req.Ukey, req.Ukey, req.Sign)
	if !verfied {
		w.WriteHeader(401)
		return
	}
	err = redis.Lock(req.Ukey)
	if err != nil {
		w.WriteHeader(423)
		return
	}
	defer redis.Unlock(req.Ukey)
	user := User{}
	err = mongo.Get("user", "ukey", req.Ukey, &user)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	err = redis.Lock(req.Mkey)
	if err != nil {
		w.WriteHeader(423)
		return
	}
	defer redis.Unlock(req.Mkey)
	order := Order{}
	err = mongo.Get2kv("trades", "ukey", req.Ukey, "mkey", req.Mkey, &order)
	if err != nil {
		w.WriteHeader(409)
		return
	}
	user.Balance += order.Offer
	if user.Balance < 0 {
		w.WriteHeader(503)
		return
	}
	mongo.Update("user", "ukey", req.Ukey, user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
}
