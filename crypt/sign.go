package crypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
)

func Sign(message string) (string, error) {
	h := sha512.New()
	h.Write([]byte(message))
	d := h.Sum(nil)
	sign, signError := rsa.SignPSS(rand.Reader, priv, crypto.SHA512, d, nil)
	if signError != nil {
		return "", signError
	}
	signBase64 := base64.StdEncoding.EncodeToString(sign)
	return signBase64, nil
}
