package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func GenerateKeysInPem() (string, error) {
	reader := rand.Reader

	key, generateKeyError := rsa.GenerateKey(reader, 4096)
	if generateKeyError != nil {
		return "", errors.New("key generation error")
	}

	privatePem := PrivateKeyToPem(key)
	publicPem := PublicKeyToPem(&key.PublicKey)
	return privatePem + "|" + publicPem, nil
}

func GenerateKeysInBytes() ([]byte, []byte) {
	reader := rand.Reader
	key, _ := rsa.GenerateKey(reader, 4096)
	privateKey := x509.MarshalPKCS1PrivateKey(key)
	publicKey := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	return privateKey, publicKey
}

func PrivateKeyToPem(key *rsa.PrivateKey) string {
	var pemKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	pemBytes := pem.EncodeToMemory(pemKey)
	return string(pemBytes)
}

func PublicKeyToPem(key *rsa.PublicKey) string {
	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(key),
	}

	pemBytes := pem.EncodeToMemory(pemkey)
	return string(pemBytes)
}

func PemToPrivateKey(pemKey string) (*rsa.PrivateKey, error) {
	privateBlock, _ := pem.Decode([]byte(pemKey))

	key, parseError := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if parseError != nil {
		return nil, errors.New("pem private parsing error")
	}
	return key, nil
}

func PemToPublicKey(pemKey string) (*rsa.PublicKey, error) {
	privateBlock, _ := pem.Decode([]byte(pemKey))

	key, parseError := x509.ParsePKCS1PublicKey(privateBlock.Bytes)
	if parseError != nil {
		return nil, errors.New("pem public parsing error")
	}
	return key, nil
}
