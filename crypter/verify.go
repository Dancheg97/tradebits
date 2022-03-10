package crypter

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
)

func (c *crypter) Verify(message string, pubkey string, sign string) error {
	pubBytes, err := base64.RawStdEncoding.DecodeString(pubkey)
	if err != nil {
		return err
	}
	signBytes, err := base64.RawStdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	pubKey, err := x509.ParsePKCS1PublicKey(pubBytes)
	if err != nil {
		return err
	}
	h := sha512.New()
	h.Write([]byte(message))
	verifyErr := rsa.VerifyPKCS1v15(
		pubKey,
		crypto.SHA512,
		h.Sum(nil),
		signBytes,
	)
	return verifyErr
}
