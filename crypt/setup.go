package crypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var priv *rsa.PrivateKey
var pub *rsa.PublicKey
var settled bool

func setPrivate(private string) error {
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

func setPublic(public string) error {
	block, _ := pem.Decode([]byte(public))
	if block == nil {
		return errors.New("failed to parse PEM block containing the key")
	}
	publickey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub = publickey
	return nil
}

func Setup(privatePEM string, publicPEM string) error {
	privErr := setPrivate(privatePEM)
	if privErr != nil {
		return privErr
	}
	return setPublic(publicPEM)
}
ad