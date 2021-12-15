package config

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
	yfile, err := ioutil.ReadFile("config2.yaml")
	if err != nil {
		log.Fatal(err)
	}
	database := map[string]string{}
	yaml.Unmarshal(yfile, &database)
	return &configuration{
		PrivatePem:       database["PrivatePem"],
		PublicPem:        database["PublicPem"],
		ConnectionAdress: database["ConnectionAdress"],
	}
}
