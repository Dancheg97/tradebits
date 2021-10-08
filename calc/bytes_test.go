package calc

import (
	"testing"
)

func TestNumToBytes(t *testing.T) {
	number := uint64(1823879123)
	bytes := NumberToBytes(number)
	if len(bytes) != 8 {
		t.Error("byte length of the number should be 8")
	}
}

func TestGenerateRandomBytes(t *testing.T) {
	bytes := Rand()
	if bytes != nil {
		return
	}
	t.Error("failed to generate random bytes")
}
