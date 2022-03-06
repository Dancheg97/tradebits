package crypter

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
)

func (c *crypter) Sign(message string) (string, error) {
	h := sha512.New()
	h.Write([]byte(message))
	sign, err := rsa.SignPKCS1v15(
		rand.Reader,
		c.priv,
		crypto.SHA512,
		h.Sum(nil),
	)
	if err != nil {
		return "", err
	}
	signBase64 := base64.RawStdEncoding.EncodeToString(sign)
	return signBase64, nil
}
