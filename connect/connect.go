package connect

import (
	"bytes"
	"tradebits/crypter"
)

type ConnectInfo struct {
	ConnectAdress string
	SelfName      string
	SelfAdress    string
}

type ConnectRequest struct {
	Hhey       string `json:"hhey"`
	Mkey       string `json:"mkey"`
	Marketname string `json:"marketname"`
	Marketlink string `json:"marketlink"`
	Sign       string `json:"sign"`
}

func Connect(crypt crypter.ICrypter, info ConnectInfo) error {
	var bytestring bytes.Buffer
	bytestring.WriteString(info.ConnectAdress)
	bytestring.WriteString(crypt.Pub())
	bytestring.WriteString(info.SelfName)
	bytestring.WriteString(info.SelfAdress)
	sign, err := crypt.Sign(bytestring.String())
	if err != nil {
		return err
	}

}
