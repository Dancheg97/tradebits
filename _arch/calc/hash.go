package calc

import "crypto/sha512"

// take hash from that byte array
func Hash(bytes []byte) []byte {
	hasher := sha512.New()
	hasher.Write(bytes)
	return hasher.Sum(nil)
}
