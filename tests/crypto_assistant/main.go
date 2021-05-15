package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"

	"golang.org/x/crypto/blake2b"
)

func Sign(message []byte, privateKey []byte) ([]byte, error) {
	private, _ := x509.ParsePKCS1PrivateKey(privateKey)
	hasher, _ := blake2b.New512([]byte("1u89hdsaj098as12"))
	hasher.Write(message)
	msgHashSum := hasher.Sum(nil)
	signatureBytes, _ := rsa.SignPSS(rand.Reader, private, crypto.BLAKE2b_512, msgHashSum, nil)
	return signatureBytes, nil
}

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

func main() {
	priv, pub := GenerateKeysInBytes()
	private := base64.RawStdEncoding.EncodeToString(priv)
	public := base64.RawStdEncoding.EncodeToString(pub)
	privFile, _ := os.Create("private.pem")
	publicFile, _ := os.Create("public.pem")
	defer privFile.Close()
	defer publicFile.Close()
	privFile.WriteString(private)
	publicFile.WriteString(public)
}
