package calc

import (
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	hash := Hash([]byte{0})
	expected := []byte{184, 36, 77, 2, 137, 129, 214, 147, 175, 123, 69, 106, 248, 239, 164, 202, 214, 61, 40, 46, 25, 255, 20, 148, 44, 36, 110, 80, 217, 53, 29, 34, 112, 74, 128, 42, 113, 195, 88, 11, 99, 112, 222, 76, 235, 41, 60, 50, 74, 132, 35, 52, 37, 87, 212, 229, 195, 132, 56, 240, 227, 105, 16, 238}
	if reflect.DeepEqual(hash, expected) {
		return
	}
	t.Error("taking blake2b hash")
}
