package _calc

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"sync_tree/__logs"
)

func Sign(message [][]byte, privateKey []byte) ([]byte, error) {
	private, privateKeyErr := x509.ParsePKCS1PrivateKey(privateKey)
	if privateKeyErr != nil {
		return nil, __logs.Error("failed to parse private key")
	}
	msgHashSum := Hash(concatenateMessage(message))
	signatureBytes, signErr := rsa.SignPSS(
		rand.Reader,
		private,
		crypto.BLAKE2b_512,
		msgHashSum,
		nil,
	)
	if signErr != nil {
		return nil, __logs.Critical("unexpected error signing message")
	}
	return signatureBytes, nil
}
