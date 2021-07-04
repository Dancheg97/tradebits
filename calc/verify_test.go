package calc

import (
	"crypto/x509"
	"io/ioutil"
	"testing"
)

func TestVerify(t *testing.T) {
	keyBytes, _ := ioutil.ReadFile("_testPriv.pem")
	mes := []byte{1, 2, 3}
	sign, _ := Sign([][]byte{mes, mes}, keyBytes)
	priv, _ := x509.ParsePKCS1PrivateKey(keyBytes)
	pubBytes := x509.MarshalPKCS1PublicKey(&priv.PublicKey)
	verified := Verify(
		[][]byte{mes, mes},
		pubBytes,
		sign,
	)
	if verified == nil {
		return
	}
	t.Error("failed to verify sign")
}

func TestBadKeyVerification(t *testing.T) {
	badKey := []byte{1, 2, 3, 4}
	sign := []byte{1, 2, 3, 4}
	mes := [][]byte{badKey, badKey}
	err := Verify(mes, badKey, sign)
	if err == nil {
		t.Error("should be an error here, cuz key bytes are invalid")
	}
}
