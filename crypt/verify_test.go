package crypt

import (
	"testing"
)

func TestVerify(t *testing.T) {
	setupForTests()
	sign, _ := Sign("")
	rez := Verify("", Pub, sign)
	if !rez {
		t.Error("sign was not verified")
	}
}
