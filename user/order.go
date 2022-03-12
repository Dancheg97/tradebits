package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

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
	Madr    string `json:"madr"`
	Mname   string `json:"mname"`
	Offer   int    `json:"offer"`
	Recieve int    `json:"recieve"`
	Sign    string `json:"sign"`
}

type Market struct {
	Name string `json:"name"`
	Mkey string `json:"mkey"`
	Link string `json:"link"`
}

func UserOrderPost(w http.ResponseWriter, r *http.Request) {
	req := OrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(406)
		return
	}
	if req.Hkey != crypt.Pub() {
		w.WriteHeader(421)
		return
	}
	var bytestring bytes.Buffer
	bytestring.WriteString(req.Hkey)
	bytestring.WriteString(req.Ukey)
	bytestring.WriteString(req.Mkey)
	bytestring.WriteString(req.Madr)
	bytestring.WriteString(req.Mname)
	bytestring.WriteString(strconv.Itoa(req.Offer))
	bytestring.WriteString(strconv.Itoa(req.Recieve))
	err = crypt.Verify(bytestring.String(), req.Ukey, req.Sign)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	err = redis.Lock(req.Ukey)
	if err != nil {
		w.WriteHeader(423)
		return
	}
	defer redis.Unlock(req.Ukey)
	err = redis.Lock(req.Mkey)
	if err != nil {
		w.WriteHeader(423)
		return
	}
	defer redis.Unlock(req.Mkey)
	user := User{}
	err = mongo.Get("user", "ukey", req.Ukey, &user)
	if err != nil {
		w.WriteHeader(503)
		return
	}
	if user.Balance < req.Offer {
		w.WriteHeader(409)
		return
	}
	m := Market{}
	err = mongo.Get("net", "mkey", req.Mkey, &m)
	if err != nil {

	}
	// TODO check trades on market foreign market
	// TODO check wether some of them could be operated
	// TODO insert trading logic
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
