package connect

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"tradebits/crypter"
	"tradebits/mongoer"
)

type ConnectInfo struct {
	ConnectAdress string
	ConnectKey    string
	SelfName      string
	SelfAdress    string
}

type ConnectRequest struct {
	Hhey  string `json:"hhey"`
	Mkey  string `json:"mkey"`
	Mname string `json:"marketname"`
	Madr  string `json:"marketlink"`
	Sign  string `json:"sign"`
}

func Connect(
	crypt crypter.ICrypter,
	mongo mongoer.IMongoer,
	info ConnectInfo,
) error {
	err := saveInformation(info.ConnectAdress, mongo)
	if err != nil {
		return err
	}
	var bytestring bytes.Buffer
	bytestring.WriteString(info.ConnectAdress)
	bytestring.WriteString(crypt.Pub())
	bytestring.WriteString(info.SelfName)
	bytestring.WriteString(info.SelfAdress)
	sign, err := crypt.Sign(bytestring.String())
	if err != nil {
		return err
	}
	request, err := json.Marshal(ConnectRequest{
		Hhey:  info.ConnectKey,
		Mkey:  crypt.Pub(),
		Mname: info.SelfName,
		Madr:  info.SelfAdress,
		Sign:  sign,
	})
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(request)
	req, err := http.NewRequest("PUT", info.ConnectAdress, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	code := resp.StatusCode
	if code != 200 {
		return errors.New("Bad status code on response " + strconv.Itoa(code))
	}
	return nil
}
