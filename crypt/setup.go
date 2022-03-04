package crypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var priv *rsa.PrivateKey

func Setup(private string) error {
	block, _ := pem.Decode([]byte(private))
	if block == nil {
		return errors.New("failed to parse PEM block containing the key")
	}
	privatekey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	priv = privatekey
	return nil
}
