package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type configuration struct {
	PrivatePem    string `yaml:"private_pem"`
	PublicPem     string `yaml:"public_pem"`
	ConnectAdress string `yaml:"connection_adress"`
}

func main() {
	yamlFile, _ := ioutil.ReadFile("config.yaml")
	conf := configuration{}

	yaml.Unmarshal(yamlFile, conf)

	fmt.Println(conf)
}
