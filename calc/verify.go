package calc

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"errors"
)

// check some sign with some public key for some and message, message
// will be concatenated to 1d byte slice
func Verify(message [][]byte, keyBytes []byte, sign []byte) error {
	publicKey, publicKeyError := x509.ParsePKCS1PublicKey(keyBytes)
	if publicKeyError != nil {
		return errors.New("error parsing public key")
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
