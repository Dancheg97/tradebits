package main

import (
	"fmt"

	"golang.org/x/crypto/blake2b"
)

func hash(key []byte) []byte {
	hasher, _ := blake2b.New512([]byte("1u89hdsaj098as12"))
	hasher.Write(key)
	hash := hasher.Sum(nil)
	return hash
}

func main() {
	randomKey := []byte{0, 1, 2, 3}
	ye := hash(randomKey)
	fmt.Println(len(ye))
}
