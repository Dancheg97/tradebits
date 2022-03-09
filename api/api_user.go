package api

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Key      string   `bson:"ukey"`
	Balance  int      `bson:"balance"`
	Messages []string `bson:"messages"`
}

type Trade struct {
	Ukey    string `bson:"ukey"`
	Mkey    string `bson:"mkey"`
	Offer   int    `bson:"offer"`
	Recieve int    `bson:"recieve"`
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
	notFound := mongo.Get("user", "ukey", ukey, &user)
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
	exists := mongo.Check("user", "ukey", ukey)
	if exists {
		w.WriteHeader(404)
		return
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserCreatePut(w http.ResponseWriter, r *http.Request) {
	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)
	hkey, exist1 := request["hkey"]
	ukey, exist2 := request["ukey"]
	sign, exist3 := request["sign"]
	if !exist1 || !exist2 || !exist3 {
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
	exists := mongo.Check("user", "ukey", ukey)
	if exists {
		w.WriteHeader(403)
		return
	}
	mongo.Put("user", User{
		Key:      ukey,
		Balance:  0,
		Messages: []string{},
	})
	w.WriteHeader(201)
}

func UserMessagePut(w http.ResponseWriter, r *http.Request) {
	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)
	hkey, exist1 := request["hkey"]
	ukey, exist2 := request["ukey"]
	mess, exist3 := request["message"]
	sign, exist4 := request["sign"]
	if !exist1 || !exist2 || !exist3 || !exist4 {
		w.WriteHeader(406)
		return
	}
	if hkey != crypt.Pub() {
		w.WriteHeader(421)
		return
	}
	verfied := crypt.Verify(hkey+ukey+mess, ukey, sign)
	if !verfied {
		w.WriteHeader(401)
		return
	}
	lockedSuccess := redis.Lock(ukey)
	if !lockedSuccess {
		w.WriteHeader(423)
		return
	}
	defer redis.Unlock(ukey)
	user := User{}
	notFound := mongo.Get("user", "ukey", ukey, &user)
	if notFound != nil {
		w.WriteHeader(404)
		return
	}
	user.Messages = append(user.Messages, mess)
	updateErr := mongo.Update("user", "ukey", ukey, user)
	if updateErr != nil {
		w.WriteHeader(503)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
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
	notFound := mongo.Get("user", "ukey", ukey, &user)
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
