package calc

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

/* This function is made to generate a pair fof pairs of priv/pub keys,
it returns 4 byte arrays for each key in that order:
 - pers priv
 - pers pub
 - mes priv
 - mes pub
*/
func Key() {
	persKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	mesKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	persPriv := x509.MarshalPKCS1PrivateKey(persKey)
	persPub := x509.MarshalPKCS1PublicKey(&persKey.PublicKey)
	
}
