package lib

import (
	"io/ioutil"
	"microservice-email/utils"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Smtp struct {
		Host     string
		Port     int
		User     string
		Password string
	}
	RabbitMQ struct {
		Host         string
		User         string
		Password     string
		QueueName    string `yaml:"queue_name"`
		ExchangeName string `yaml:"exchange_name"`
		ExchangeKind string `yaml:"exchange_kind"`
		Declare      bool
	}
}

var Conf *Config
var ConfigFilePath string

func ReadConfig() {
	fileData, err := ioutil.ReadFile(ConfigFilePath)
	utils.CheckException(err)

	Conf = &Config{}

	err = yaml.Unmarshal(fileData, Conf)
	utils.CheckException(err)
}
