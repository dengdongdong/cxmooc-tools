package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	MySQL MySQL
	Redis Redis
	VCode string
}

type Redis struct {
	Addr     string
	Password string
	DB       int
}

type MySQL struct {
	Dsn string
}

var AppConfig Config

func init() {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("config read error: %v", err)
	}
	err = yaml.Unmarshal(file, AppConfig)
	if err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}
}
