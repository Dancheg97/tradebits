package _calc

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"sync_tree/__logs"
)

func Sign(message [][]byte, privateKey []byte) ([]byte, error) {
	private, privateKeyErr := x509.ParsePKCS1PrivateKey(privateKey)
	if privateKeyErr != nil {
		__logs.Error(errors.New("failed to parse private key"))
		return
	}
	msgHashSum := Hash(concatenateMessage(message))
	signatureBytes, signErr := rsa.SignPSS(
		rand.Reader,
		private,
		crypto.BLAKE2b_512,
		msgHashSum,
		nil,
	)
	if hasherErr != nil {
		__logs.Critical(errors.New("unexpected error signing message"))
		return
	}
	return signatureBytes, nil
}
