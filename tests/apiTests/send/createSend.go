package main

import (
	"bc_server/calc"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/blake2b"
)

func Sign(message [][]byte, privateKey []byte) ([]byte, error) {
	mes := ConcatenateMessage(message)
	private, _ := x509.ParsePKCS1PrivateKey(privateKey)
	hasher, _ := blake2b.New512([]byte("1u89hdsaj098as12"))
	hasher.Write(mes)
	msgHashSum := hasher.Sum(nil)
	signatureBytes, _ := rsa.SignPSS(rand.Reader, private, crypto.BLAKE2b_512, msgHashSum, nil)
	return signatureBytes, nil
}

func ConcatenateMessage(message [][]byte) []byte {
	concatenated := []byte{}
	for i := 0; i < len(message); i++ {
		concatenated = append(concatenated, message[i]...)
	}
	return concatenated
}

type sendRequest struct {
	SenderPublicKey []byte `json:"SenderPublicKey"`
	SendAmountBytes []byte `json:"SendAmountBytes"`
	RecieverAdress  []byte `json:"RecieverAdress"`
	SenderSign      []byte `json:"SenderSign"`
}

func Verify(message [][]byte, keyBytes []byte, sign []byte) error {
	publicKey, publicKeyError := x509.ParsePKCS1PublicKey(keyBytes)
	if publicKeyError != nil {
		return errors.New("public key error")
	}
	hasher, _ := blake2b.New512([]byte("1u89hdsaj098as12"))
	hasher.Write([]byte(ConcatenateMessage(message)))
	hash := hasher.Sum(nil)
	return rsa.VerifyPSS(publicKey, crypto.BLAKE2b_512, hash, sign, nil)
}

func main() {
	publicBytesAsStringBase64, _ := ioutil.ReadFile("public.pem")
	publicBytes, _ := base64.RawStdEncoding.DecodeString(string(publicBytesAsStringBase64))
	privateBytesAsStringBase64, _ := ioutil.ReadFile("private.pem")
	privateBytes, _ := base64.RawStdEncoding.DecodeString(string(privateBytesAsStringBase64))
	recieverBytesAsStringBase64, _ := ioutil.ReadFile("public2.pem")
	recieverBytes, _ := base64.RawStdEncoding.DecodeString(string(recieverBytesAsStringBase64))
	amount := make([]byte, 8)
	binary.LittleEndian.PutUint64(amount, uint64(1500))
	recieverAdress := calc.Hash(recieverBytes)
	message := [][]byte{publicBytes, amount, recieverAdress}
	sign, _ := Sign(message, privateBytes)
	request := sendRequest{
		SenderPublicKey: publicBytes,
		SendAmountBytes: amount,
		RecieverAdress:  recieverAdress,
		SenderSign:      sign,
	}
	req, _ := json.Marshal(request)
	fmt.Println(string(req))
}