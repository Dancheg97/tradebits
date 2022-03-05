package crypt

import (
	"testing"
)

func TestSign(t *testing.T) {
	setupForTests()
	sign, err := Sign("")
	if err != nil {
		t.Error(err)
	}
	if len(sign) < 10 {
		t.Error("bad sign lengt")
	}
}
