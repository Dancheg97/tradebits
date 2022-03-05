package crypt

import (
	"testing"
)

func TestVerify(t *testing.T) {
	setupForTests()
	sign, _ := Sign("")
	signVerificationSuccess := Verify("", Pub, sign)
	if !signVerificationSuccess {
		t.Error("sign was not verified")
	}
}
