package calc

import (
	"crypto/x509"
	"io/ioutil"
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

func TestSign(t *testing.T) {
	keyBytes, _ := ioutil.ReadFile("_testPriv.pem")
	mes := []byte{1, 2, 3}
	sign, _ := Sign([][]byte{mes, mes}, keyBytes)
	if len(sign) == 256 {
		return
	}
	t.Error("signature error")
}

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

func TestGenerateRandomBytes(t *testing.T) {
	bytes := Rand()
	if bytes != nil {
		return
	}
	t.Error("failed to generate random bytes")
}

func TestGetKeys(t *testing.T) {
	keys := Gen()
	lenSum := len(keys.PersPriv) + len(keys.PersPub) + len(keys.MesPriv) + len(keys.MesPub)
	if lenSum < 4330 || lenSum > 4340 {
		t.Error("failed to generate correct keys")
	}
}

func TestNumToBytes(t *testing.T) {
	number := uint64(1823879123)
	bytes := NumberToBytes(number)
	if len(bytes) != 8 {
		t.Error("byte length of the number should be 8")
	}
}

func TestBadKeySign(t* testing.T) {
	badKey := []byte{1, 2, 3, 4}
	mes := [][]byte{badKey, badKey}
	_, err := Sign(mes, badKey)
	if err == nil {
		t.Error("should get an error cuz key is invalid")
	}
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