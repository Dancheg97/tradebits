package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"strings"
)

type configuration struct {
	PrivatePem       []byte
	PublicPem        []byte
	ConnectionAdress string
}

func readKeyBytes(filename string) []byte {
	privBytes, fileErr := ioutil.ReadFile(filename)
	if fileErr != nil {
		log.Fatal("key file not found")
	}
	block, _ := pem.Decode(privBytes)
	_, privErr := x509.ParsePKCS1PrivateKey(block.Bytes)
	_, pubErr := x509.ParsePKCS1PublicKey(block.Bytes)
	if pubErr != nil && privErr != nil {
		log.Fatal(privErr, pubErr)
	}
	return block.Bytes
}

func readConnecitonAdress(filename string) string {
	adressBytes, fileErr := ioutil.ReadFile(filename)
	if fileErr != nil {
		log.Fatal("connection file not found")
	}
	if !strings.Contains(string(adressBytes), ".") {
		log.Fatal("adress has errors")
	}
	return string(adressBytes)
}

// this function is made to read configuration file, if something is missing
// will raise fatal error
func ReadConfiguration() *configuration {
	return &configuration{
		PrivatePem:       readKeyBytes(".private.pem"),
		PublicPem:        readKeyBytes(".public.pem"),
		ConnectionAdress: readConnecitonAdress(".connect.cfg"),
	}
}
