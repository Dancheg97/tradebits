package calc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

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
