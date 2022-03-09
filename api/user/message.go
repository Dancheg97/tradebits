package user

import (
	"encoding/json"
	"net/http"
)

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
