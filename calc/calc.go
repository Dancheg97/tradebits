package calc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/binary"
)

// take hash from that byte array
func Hash(bytes []byte) []byte {
	hasher := sha512.New()
	hasher.Write(bytes)
	return hasher.Sum(nil)
}

type Keys struct {
	PersPriv []byte
	PersPub  []byte
	MesPriv  []byte
	MesPub   []byte
}

/* This function is made to generate a pair fof pairs of priv/pub keys,
it returns 4 byte arrays for each key in that order:
 - pers priv
 - pers pub
 - mes priv
 - mes pub
*/
func Gen() Keys {
	persKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	mesKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return Keys{
		PersPriv: x509.MarshalPKCS1PrivateKey(persKey),
		PersPub:  x509.MarshalPKCS1PublicKey(&persKey.PublicKey),
		MesPriv:  x509.MarshalPKCS1PrivateKey(mesKey),
		MesPub:   x509.MarshalPKCS1PublicKey(&mesKey.PublicKey),
	}
}

func NumberToBytes(number uint64) []byte {
	amountBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountBytes, number)
	return amountBytes
}

func Rand() []byte {
	randomBytes := make([]byte, 64)
	rand.Read(randomBytes)
	return randomBytes
}
