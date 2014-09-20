package main

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	MysqlHost string
	MysqlUser string
	MysqlPass string
	MysqlDb   string

	CloudflareKey   string
	CloudflareEmail string
}

var config *Configuration

func getConfig() *Configuration {
	if config == nil {
		config = loadConfig("config.json")
	}

	return config
}

func loadConfig(path string) *Configuration {
	configFile, err := os.Open(path)

	if err != nil {
		log.Fatalf("Error while trying to read %s: %s", path, err)
	}

	decoder := json.NewDecoder(configFile)

	c := Configuration{}
	err = decoder.Decode(&c)

	if err != nil {
		log.Fatalf("Error parsing config file: %s", err)
	}

	return &c
}
