package user

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Key      string   `bson:"ukey"`
	Balance  int      `bson:"balance"`
	Messages []string `bson:"messages"`
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
