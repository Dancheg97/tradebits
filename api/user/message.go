package user

import (
	"encoding/json"
	"net/http"
)

type MessageRequest struct {
	Hkey string `json:"hkey"`
	Ukey string `json:"ukey"`
	Mess string `json:"mess"`
	Sign string `json:"sign"`
}

func UserMessagePut(w http.ResponseWriter, r *http.Request) {
	req := MessageRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(406)
		return
	}
	if req.Hkey != crypt.Pub() {
		w.WriteHeader(421)
		return
	}
	verfied := crypt.Verify(req.Hkey+req.Ukey+req.Mess, req.Ukey, req.Sign)
	if !verfied {
		w.WriteHeader(401)
		return
	}
	lockedSuccess := redis.Lock(req.Ukey)
	if !lockedSuccess {
		w.WriteHeader(423)
		return
	}
	defer redis.Unlock(req.Ukey)
	user := User{}
	notFound := mongo.Get("user", "ukey", req.Ukey, &user)
	if notFound != nil {
		w.WriteHeader(404)
		return
	}
	user.Messages = append(user.Messages, req.Mess)
	updateErr := mongo.Update("user", "ukey", req.Ukey, user)
	if updateErr != nil {
		w.WriteHeader(503)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
}
