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
	sign, err := rsa.SignPSS(
		rand.Reader,
		priv,
		crypto.SHA512,
		h.Sum(nil),
		nil,
	)
	if err != nil {
		return "", err
	}
	signBase64 := base64.StdEncoding.EncodeToString(sign)
	return signBase64, nil
}