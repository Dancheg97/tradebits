package crypter

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto_assistant/keys"
	"errors"
)

func Encrypt(message string, publicPem string) (string, error) {
	publicKey, keyParseError := keys.PemToPublicKey(publicPem)
	if keyParseError != nil {
		return "", errors.New("error parsing public key")
	}

	messageAsBytes := []byte(message)
	encryptedMessage, encryptionError := rsa.EncryptPKCS1v15(rand.Reader, publicKey, messageAsBytes)
	if encryptionError != nil {
		return "", errors.New("error encrypting message")
	}

	return string(encryptedMessage), nil
}

func Decrypt(message string, privatePem string) (string, error) {
	privateKey, keyParseError := keys.PemToPrivateKey(privatePem)
	if keyParseError != nil {
		return "", errors.New("error parsing private key")
	}

	messageAsBytes := []byte(message)
	decryptedMessage, decryptionError := rsa.DecryptPKCS1v15(rand.Reader, privateKey, messageAsBytes)
	if decryptionError != nil {
		return "", errors.New("error decrypting message")
	}

	return string(decryptedMessage), nil
}
