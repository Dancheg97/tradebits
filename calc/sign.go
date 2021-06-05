package calc

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
)

// sign message with private key (message is taken by 2d ar, and will be
// cncatenated to single one before signing)
func Sign(message [][]byte, privateKey []byte) ([]byte, error) {
	private, privateKeyErr := x509.ParsePKCS1PrivateKey(privateKey)
	if privateKeyErr != nil {
		return nil, errors.New("parse private key error")
	}
	msgHashSum := Hash(concatenateMessage(message))
	signatureBytes, _ := rsa.SignPSS(
		rand.Reader,
		private,
		crypto.BLAKE2b_512,
		msgHashSum,
		nil,
	)
	return signatureBytes, nil
}
