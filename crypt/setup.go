package crypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

var priv *rsa.PrivateKey
var Pub string

func Setup(privateBase64 string) error {
	keyBytes, decodeErr := base64.RawStdEncoding.DecodeString(privateBase64)
	if decodeErr != nil {
		return decodeErr
	}
	privatekey, parseErr := x509.ParsePKCS1PrivateKey(keyBytes)
	if parseErr != nil {
		return parseErr
	}
	priv = privatekey
	pubBytes := x509.MarshalPKCS1PublicKey(&privatekey.PublicKey)
	Pub = base64.RawStdEncoding.EncodeToString(pubBytes)
	return nil
}
