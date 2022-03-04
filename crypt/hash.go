package crypt

import (
	"crypto/sha512"
)

func Hash(bytes []byte) []byte {
	hasher := sha512.New()
	hasher.Write(bytes)
	return hasher.Sum(nil)
}
