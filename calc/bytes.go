package calc

import (
	"crypto/rand"
	"encoding/binary"
)

func NumberToBytes(number uint64) []byte {
	amountBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountBytes, number)
	return amountBytes
}

func Rand() []byte {
	randomBytes := make([]byte, 64)
	rand.Read(randomBytes)
	return randomBytes
}
