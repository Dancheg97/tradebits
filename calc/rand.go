package calc

import (
	"crypto/rand"
	"sync_tree/logs"
)

func Rand() []byte {
	randomBytes := make([]byte, 64)
	_, randomGenerateErr := rand.Read(randomBytes)
	if randomGenerateErr != nil {
		logs.Critical("error generating random bytes for market")
		return nil
	}
	return randomBytes
}
