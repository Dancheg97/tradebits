package calc

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)


func TestHash(t *testing.T) {
	hash := Hash([]byte{0})
	fmt.Println(hash)
	expected := []byte{82, 183, 9, 232, 198, 142, 69, 166, 187, 232, 192, 96, 68, 75, 73, 151, 112, 183, 61, 164, 253, 193, 184, 5, 181, 233, 156, 208, 175, 166, 26, 164, 209, 27, 80, 16, 145, 0, 78, 104, 11, 151, 150, 131, 100, 121, 160, 43, 34, 196, 98, 251, 92, 80, 223, 142, 20, 163, 21, 232, 171, 85, 72, 39}
	if reflect.DeepEqual(hash, expected) {
		return
	}
	t.Error("taking blake2b hash")
}

func TestSign(t *testing.T) {
	keyBytes, _ := ioutil.ReadFile("calc/priv.pem")
	mes := []byte{1, 2, 3}
	sign, _ := Sign([][]byte{mes, mes}, keyBytes)
	if len(sign) == 512 {
		return
	}
	t.Error("taking blake2b hash")
}

func TestVerify(t *testing.T) {
	keyBytes, _ := ioutil.ReadFile("calc/priv.pem")
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
	if len(keys.MesPriv) != 1190 || len(keys.MesPriv) != 1192 {
		t.Error("Wrong mes priv key length", len(keys.MesPriv))
	}
	if len(keys.MesPub) != 270 {
		t.Error("Wrong mes pub key ltngth")
	}
	if len(keys.PersPriv) != 2348 {
		t.Error("Wrong pers priv key length")
	}
	if len(keys.PersPub) != 526 {
		t.Error("Wrong pers pub key length")
	}
}
