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
