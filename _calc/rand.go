package _calc

import (
	"crypto/rand"
	"sync_tree/__logs"
)

func RandBytes() []byte {
	randomBytes := make([]byte, 64)
	_, randomGenerateErr := rand.Read(randomBytes)
	if randomGenerateErr != nil {
		__logs.Critical("error generating random bytes for asset")
		return nil
	}
	return randomBytes
}
