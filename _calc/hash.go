package _calc

import (
	"errors"
	"sync_tree/__logs"

	"golang.org/x/crypto/blake2b"
)

func Hash(key []byte) []byte {
	hasher, hashErr := blake2b.New512([]byte("1u89hdsaj098as12"))
	if hashErr != nil {
		__logs.Critical(errors.New("unexpected error craeting hasher"))
		return
	}
	hasher.Write(key)
	hash := hasher.Sum(nil)
	return hash
}
