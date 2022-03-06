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
	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)
	ukey, exists := request["ukey"]
	if !exists {
		w.WriteHeader(406)
		return
	}
	user := User{}
	notFound := mongo.Get(ukey, "user", &user)
	if notFound != nil {
		w.WriteHeader(404)
		return
	}
	response := map[string]int{
		"balance": user.Balance,
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

func UserCancelordersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserCreatePut(w http.ResponseWriter, r *http.Request) {
	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)
	hkey, exist1 := request["hkey"]
	ukey, exist2 := request["ukey"]
	sign, exist3 := request["sign"]
	if !(exist1 && exist2 && exist3) {
		w.WriteHeader(406)
		return
	}
	if hkey != crypt.Pub {
		w.WriteHeader(421)
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
	mongo.Put("user", User{
		Key:      ukey,
		Messages: []string{},
	})
	w.WriteHeader(201)
}

func UserMessagePut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserMessagesGet(w http.ResponseWriter, r *http.Request) {
	request := map[string]interface{}{}
	json.NewDecoder(r.Body).Decode(&request)
	iukey, exist1 := request["ukey"]
	ukey, asserted1 := iukey.(string)
	ioffset, exist2 := request["offset"]
	offsetFloat, asserted2 := ioffset.(float64)
	offset := int(offsetFloat)
	if !(exist1 && asserted1 && exist2 && asserted2) {
		w.WriteHeader(406)
		return
	}
	user := User{}
	notFound := mongo.Get(ukey, "user", &user)
	if notFound != nil {
		w.WriteHeader(404)
		return
	}
	if len(user.Messages) < offset {
		w.WriteHeader(406)
		return
	}
	response := user.Messages[offset:]
	respbytes, marshErr := json.Marshal(response)
	if marshErr != nil {
		w.WriteHeader(503)
		return
	}
	w.Write(respbytes)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
}

func UserOrderPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserTradesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
