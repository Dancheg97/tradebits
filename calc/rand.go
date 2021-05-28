package calc

import (
	"crypto/rand"
	"sync_tree/logs"
)

func RandBytes() []byte {
	randomBytes := make([]byte, 64)
	_, randomGenerateErr := rand.Read(randomBytes)
	if randomGenerateErr != nil {
		logs.Critical("error generating random bytes for asset")
		return nil
	}
	return randomBytes
}
