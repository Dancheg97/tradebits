package calc

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"sync_tree/logs"
)

// sign message with private key (message is taken by 2d ar, and will be
// cncatenated to single one before signing)
func Sign(message [][]byte, privateKey []byte) ([]byte, error) {
	private, privateKeyErr := x509.ParsePKCS1PrivateKey(privateKey)
	if privateKeyErr != nil {
		return nil, logs.Error("failed to parse private key")
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
		return nil, logs.Critical("unexpected error signing message")
	}
	return signatureBytes, nil
}
