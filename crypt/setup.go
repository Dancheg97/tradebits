package crypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

type crypter struct {
	priv *rsa.PrivateKey
	pub  string
}

func Crypter(privateBase64 string) (*crypter, error) {
	keyBytes, decodeErr := base64.RawStdEncoding.DecodeString(privateBase64)
	if decodeErr != nil {
		return nil, decodeErr
	}
	privatekey, parseErr := x509.ParsePKCS1PrivateKey(keyBytes)
	if parseErr != nil {
		return nil, parseErr
	}
	pubBytes := x509.MarshalPKCS1PublicKey(&privatekey.PublicKey)
	pubBase64 := base64.RawStdEncoding.EncodeToString(pubBytes)
	c := crypter{
		priv: privatekey,
		pub:  pubBase64,
	}
	return &c, nil
}
