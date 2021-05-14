package hash

import (
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/blake2b"
)

func GetHash(message string) (string, error) {
	hasher, hashCreationError := blake2b.New512([]byte("1u89hdsaj098as12"))
	if hashCreationError != nil {
		return "", errors.New("hasher creation error")
	}
	hasher.Write([]byte(message))
	msgHashSum := hasher.Sum(nil)

	signatureBase64 := base64.RawStdEncoding.EncodeToString(msgHashSum)
	return signatureBase64, nil
}
