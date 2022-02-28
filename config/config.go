package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type config struct {
	Serve  string `json:"serve"`
	Dgraph string `json:"dgraph"`
	Redis  string `json:"redis"`
}

func GetConfig(location string) config {
	file, fileErr := ioutil.ReadFile(location)
	if fileErr != nil {
		log.Panic("Config not found")
	}
	conf := config{}
	parseErr := json.Unmarshal([]byte(file), &conf)
	if parseErr != nil {
		log.Panic("Bad config file")
	}
	return conf
}
