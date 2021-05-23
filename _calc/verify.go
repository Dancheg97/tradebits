package _calc

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"sync_tree/__logs"
)

func Verify(message [][]byte, keyBytes []byte, sign []byte) error {
	publicKey, publicKeyError := x509.ParsePKCS1PublicKey(keyBytes)
	if publicKeyError != nil {
		return __logs.Error("error parsing public key while signing")
	}
	hash := Hash(concatenateMessage(message))
	return rsa.VerifyPSS(publicKey, crypto.BLAKE2b_512, hash, sign, nil)
}

func concatenateMessage(message [][]byte) []byte {
	concatenated := []byte{}
	for i := 0; i < len(message); i++ {
		concatenated = append(concatenated, message[i]...)
	}
	return concatenated
}
