package calc

import (
	"io/ioutil"
	"testing"
)

func TestSign(t *testing.T) {
	keyBytes, _ := ioutil.ReadFile("_testPriv.pem")
	mes := []byte{1, 2, 3}
	sign, _ := Sign([][]byte{mes, mes}, keyBytes)
	if len(sign) == 256 {
		return
	}
	t.Error("signature error")
}

func TestBadKeySign(t *testing.T) {
	badKey := []byte{1, 2, 3, 4}
	mes := [][]byte{badKey, badKey}
	_, err := Sign(mes, badKey)
	if err == nil {
		t.Error("should get an error cuz key is invalid")
	}
}
