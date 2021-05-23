package _calc

import (
	"sync_tree/__logs"

	"golang.org/x/crypto/blake2b"
)

func Hash(bytes []byte) []byte {
	hasher, hashErr := blake2b.New512([]byte("1u89hdsaj098as12"))
	if hashErr != nil {
		__logs.Critical("unexpected error craeting hasher")
		return nil
	}
	hasher.Write(bytes)
	hash := hasher.Sum(nil)
	return hash
}
