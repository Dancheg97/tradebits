package calc

import (
	"sync_tree/logs"

	"golang.org/x/crypto/blake2b"
)

// take hash from that byte array
func Hash(bytes []byte) []byte {
	hasher, hashErr := blake2b.New512([]byte("1u89hdsaj098as12"))
	if hashErr != nil {
		logs.Critical("unexpected error craeting hasher")
		return nil
	}
	hasher.Write(bytes)
	hash := hasher.Sum(nil)
	return hash
}
