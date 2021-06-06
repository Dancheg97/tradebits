package calc

import (
	"crypto/rand"
)

func Rand() []byte {
	randomBytes := make([]byte, 64)
	rand.Read(randomBytes)
	return randomBytes
}
