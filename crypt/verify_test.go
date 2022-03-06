package crypt

import (
	"testing"
)

func TestVerify(t *testing.T) {
	setupForTests()
	success := Verify(testmes, Pub, testsign)
	if !success {
		t.Error("sign was not verified")
	}
}
