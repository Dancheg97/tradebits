package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto_assistant/keys"
	"encoding/base64"
	"errors"

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

func Verify(message []byte, keyBytes []byte, sign []byte) error {
	publicKey, publicKeyError := x509.ParsePKCS1PublicKey(keyBytes)
	if publicKeyError != nil {
		return errors.New("public key error")
	}
	hasher, _ := blake2b.New512([]byte("1u89hdsaj098as12"))
	hasher.Write([]byte(message))
	hash := hasher.Sum(nil)
	return rsa.VerifyPSS(publicKey, crypto.BLAKE2b_512, hash, sign, nil)
}

func SignString(message string, privatePem string) (string, error) {
	private, pemError := keys.PemToPrivateKey(privatePem)
	if pemError != nil {
		return "", errors.New("private key parse error in signing")
	}

	hasher, hashCreationError := blake2b.New512([]byte("1u89hdsaj098as12"))
	if hashCreationError != nil {
		return "", errors.New("hasher creation error")
	}
	hasher.Write([]byte(message))
	msgHashSum := hasher.Sum(nil)

	signatureBytes, signError := rsa.SignPSS(rand.Reader, private, crypto.BLAKE2b_512, msgHashSum, nil)
	if signError != nil {
		return "", errors.New("signature creation error")
	}

	signatureBase64 := base64.RawStdEncoding.EncodeToString(signatureBytes)
	return signatureBase64, nil
}

func VerifyString(message string, pubPem string, sign string) (bool, error) {
	public, pemError := keys.PemToPublicKey(pubPem)
	if pemError != nil {
		return false, errors.New("private key parse error in signing")
	}

	hasher, hashCreationError := blake2b.New512([]byte("1u89hdsaj098as12"))
	if hashCreationError != nil {
		return false, errors.New("hasher creation error")
	}
	hasher.Write([]byte(message))
	msgHashSum := hasher.Sum(nil)

	signature, signDecodeError := base64.RawStdEncoding.DecodeString(sign)
	if signDecodeError != nil {
		return false, errors.New("sign base64 decode error")
	}

	verified := rsa.VerifyPSS(public, crypto.BLAKE2b_512, msgHashSum, signature, nil)
	if verified != nil {
		return false, errors.New("signature verification error")
	}
	return true, nil
}
