package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestSetup(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)
	rez := Setup(string(keyPEM))
	if rez != nil {
		t.Error(rez)
	}
}
