package calc

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	"golang.org/x/crypto/blake2b"
)

/*
Кулон содержит две печати
И три скрижали изнутри
Проверка правды на твоей цитате
Пройдет лишь тот, кто чист, увы
*/

func Verify(message [][]byte, keyBytes []byte, sign []byte) error {
	publicKey, publicKeyError := x509.ParsePKCS1PublicKey(keyBytes)
	if publicKeyError != nil {
		return errors.New("public key error")
	}
	hasher, _ := blake2b.New512([]byte("1u89hdsaj098as12"))
	hasher.Write([]byte(ConcatenateMessage(message)))
	hash := hasher.Sum(nil)
	return rsa.VerifyPSS(publicKey, crypto.BLAKE2b_512, hash, sign, nil)
}

func ConcatenateMessage(message [][]byte) []byte {
	concatenated := []byte{}
	for i := 0; i < len(message); i++ {
		concatenated = append(concatenated, message[i]...)
	}
	return concatenated
}

func PemToPublicKey(pemKey string) (*rsa.PublicKey, error) {
	privateBlock, _ := pem.Decode([]byte(pemKey))

	key, parseError := x509.ParsePKCS1PublicKey(privateBlock.Bytes)
	if parseError != nil {
		return nil, errors.New("pem public parsing error")
	}
	return key, nil
}

func VerifyString(message string, pubPem string, sign string) error {
	public, pemError := PemToPublicKey(pubPem)
	if pemError != nil {
		return errors.New("private key parse error in signing")
	}

	hasher, hashCreationError := blake2b.New512([]byte("1u89hdsaj098as12"))
	if hashCreationError != nil {
		return errors.New("hasher creation error")
	}
	hasher.Write([]byte(message))
	msgHashSum := hasher.Sum(nil)

	signature, signDecodeError := base64.RawStdEncoding.DecodeString(sign)
	if signDecodeError != nil {
		return errors.New("sign base64 decode error")
	}

	verified := rsa.VerifyPSS(public, crypto.BLAKE2b_512, msgHashSum, signature, nil)
	if verified != nil {
		return errors.New("signature verification error")
	}
	return nil
}
