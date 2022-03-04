package crypt

import (
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	inp := Hash([]byte{})
	expected := []byte{207, 131, 225, 53, 126, 239, 184, 189, 241, 84, 40, 80, 214, 109, 128, 7, 214, 32, 228, 5, 11, 87, 21, 220, 131, 244, 169, 33, 211, 108, 233, 206, 71, 208, 209, 60, 93, 133, 242, 176, 255, 131, 24, 210, 135, 126, 236, 47, 99, 185, 49, 189, 71, 65, 122, 129, 165, 56, 50, 122, 249, 39, 218, 62}
	if !reflect.DeepEqual(inp, expected) {
		t.Error("hashes are not matching")
	}
}
