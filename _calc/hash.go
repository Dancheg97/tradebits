package calc

import (
	"golang.org/x/crypto/blake2b"
)

func Hash(key []byte) []byte {
	hasher, _ := blake2b.New512([]byte("1u89hdsaj098as12"))
	hasher.Write(key)
	hash := hasher.Sum(nil)
	return hash
}

func HashKeyString(key string) []byte {
	hasher, _ := blake2b.New512([]byte("1u89hdsaj098as12"))
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)
	return hash
}
