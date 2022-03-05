package api

import (
	"encoding/json"
	"net/http"
	"tradebits/crypt"
	"tradebits/mongo"
)

type User struct {
	Key      string   `bson:"key"`
	Balance  int      `bson:"balance"`
	Messages []string `bson:"messages"`
}

func UserBalanceGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserCancelordersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserCreatePut(w http.ResponseWriter, r *http.Request) {
	inp := map[string]string{}
	json.NewDecoder(r.Body).Decode(&inp)
	hkey, e1 := inp["hkey"]
	ukey, e2 := inp["ukey"]
	sign, e3 := inp["sign"]
	if !(e1 && e2 && e3 && hkey == crypt.Pub) {
		w.WriteHeader(406)
		return
	}
	verfied := crypt.Verify(hkey+ukey, ukey, sign)
	if !verfied {
		w.WriteHeader(401)
		return
	}
	exists := mongo.Check(ukey, "user")
	if exists {
		w.WriteHeader(403)
		return
	}
	mongo.Put("user", User{Key: ukey})
	w.WriteHeader(http.StatusOK)
}

func UserMessagePut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserMessagesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserOrderPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserTradesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
