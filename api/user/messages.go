package user

import (
	"encoding/json"
	"net/http"
)

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
