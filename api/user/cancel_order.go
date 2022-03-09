package user

import (
	"encoding/json"
	"net/http"
)

func UserCancelordersPost(w http.ResponseWriter, r *http.Request) {
	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)
	hkey, exist1 := request["hkey"]
	ukey, exist2 := request["ukey"]
	mkey, exist3 := request["mkey"]
	sign, exist4 := request["sign"]
	if !exist1 || !exist2 || !exist3 || !exist4 {
		w.WriteHeader(406)
		return
	}
	if hkey != crypt.Pub() {
		w.WriteHeader(421)
		return
	}
	verfied := crypt.Verify(hkey+ukey, ukey, sign)
	if !verfied {
		w.WriteHeader(401)
		return
	}
	userLocked := redis.Lock(ukey)
	if !userLocked {
		w.WriteHeader(423)
		return
	}
	defer redis.Unlock(ukey)
	user := User{}
	err := mongo.Get("user", "ukey", ukey, &user)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	marketLocked := redis.Lock(mkey)
	if !marketLocked {
		w.WriteHeader(423)
		return
	}
	defer redis.Unlock(mkey)
	order := Order{}
	err = mongo.Get2kv("trades", "ukey", ukey, "mkey", mkey, &order)
	if err != nil {
		w.WriteHeader(409)
		return
	}
	user.Balance += order.Offer
	if user.Balance < 0 {
		w.WriteHeader(503)
		return
	}
	mongo.Update("user", "ukey", ukey, user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
}
