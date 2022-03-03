package crypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func Sign(base64str string) (string, error) {
	bytes, decodeError := base64.StdEncoding.DecodeString(base64str)
	if decodeError != nil {
		return "", decodeError
	}
	h := sha256.New()
	h.Write(bytes)
	d := h.Sum(nil)
	sign, signError := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, d)
	if signError != nil {
		return "", signError
	}
	signBase64 := base64.StdEncoding.EncodeToString(sign)
	return signBase64, nil
}
