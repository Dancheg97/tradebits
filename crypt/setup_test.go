package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestSetup(t *testing.T) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 64)
	privBytes := x509.MarshalPKCS1PrivateKey(privkey)
	privPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privBytes,
		},
	)
	rez := Setup(string(privPem))
	if rez != nil {
		t.Error(rez)
	}

}
