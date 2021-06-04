package calc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

type keys struct {
	persPriv []byte
	persPub  []byte
	mesPriv  []byte
	mesPub   []byte
}

/* This function is made to generate a pair fof pairs of priv/pub keys,
it returns 4 byte arrays for each key in that order:
 - pers priv
 - pers pub
 - mes priv
 - mes pub
*/
func Gen() *keys {
	persKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	mesKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return &keys{
		persPriv: x509.MarshalPKCS1PrivateKey(persKey),
		persPub:  x509.MarshalPKCS1PublicKey(&persKey.PublicKey),
		mesPriv:  x509.MarshalPKCS1PrivateKey(mesKey),
		mesPub:   x509.MarshalPKCS1PublicKey(&mesKey.PublicKey),
	}
}
