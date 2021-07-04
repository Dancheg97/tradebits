package calc

import (
	"testing"
)

func TestGenerateRandomBytes(t *testing.T) {
	bytes := Rand()
	if bytes != nil {
		return
	}
	t.Error("failed to generate random bytes")
}
