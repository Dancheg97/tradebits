package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type configuration struct {
	PrivatePem       string
	PublicPem        string
	ConnectionAdress string
}

func GetConfiguration() *configuration {
	privBytes, err := ioutil.ReadFile(".private.pem")
	if err != nil {
		log.Fatal("private key file not found")
	}
	publicBytes, err := ioutil.ReadFile(".public.pem")
	if err != nil {
		log.Fatal("public key file not found")
	}
	connectionBytes, err := ioutil.ReadFile(".connect.cfg")
	if err != nil {
		log.Fatal("connection file not found")
	}
}
