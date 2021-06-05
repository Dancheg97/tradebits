package calc

import (
	"golang.org/x/crypto/blake2b"
)

// take hash from that byte array
func Hash(bytes []byte) []byte {
	hasher, _ := blake2b.New512([]byte("1u89hdsaj098as12"))
	hasher.Write(bytes)
	hash := hasher.Sum(nil)
	return hash
}
