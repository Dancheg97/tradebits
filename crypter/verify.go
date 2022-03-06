package crypt

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
)

func (c *crypter) Verify(message string, pubkey string, sign string) bool {
	pubBytes, decodeKeyErr := base64.RawStdEncoding.DecodeString(pubkey)
	if decodeKeyErr != nil {
		return false
	}
	signBytes, decodeSignErr := base64.RawStdEncoding.DecodeString(sign)
	if decodeSignErr != nil {
		return false
	}
	pubKey, parseErr := x509.ParsePKCS1PublicKey(pubBytes)
	if parseErr != nil {
		return false
	}
	h := sha512.New()
	h.Write([]byte(message))
	verifyErr := rsa.VerifyPKCS1v15(
		pubKey,
		crypto.SHA512,
		h.Sum(nil),
		signBytes,
	)
	return verifyErr == nil
}
