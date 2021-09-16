package calc

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/binary"
	"errors"
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

// sign message with private key (message is taken by 2d ar, and will be
// cncatenated to single one before signing)
func Sign(message [][]byte, privateKey []byte) ([]byte, error) {
	private, privateKeyErr := x509.ParsePKCS1PrivateKey(privateKey)
	if privateKeyErr != nil {
		return nil, errors.New("parse private key error")
	}
	msgHashSum := Hash(concatenateMessage(message))
	signatureBytes, _ := rsa.SignPKCS1v15(
		rand.Reader,
		private,
		crypto.SHA512,
		msgHashSum,
	)
	return signatureBytes, nil
}

// check some sign with some public key for some and message, message
// will be concatenated to 1d byte slice
func Verify(message [][]byte, keyBytes []byte, sign []byte) error {
	publicKey, publicKeyError := x509.ParsePKCS1PublicKey(keyBytes)
	if publicKeyError != nil {
		return errors.New("error parsing public key")
	}
	hash := Hash(concatenateMessage(message))
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hash, sign)
}

func concatenateMessage(message [][]byte) []byte {
	concatenated := []byte{}
	for i := 0; i < len(message); i++ {
		concatenated = append(concatenated, message[i]...)
	}
	return concatenated
}
