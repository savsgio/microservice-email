package lib

import (
	"flag"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"microservice-email/utils"
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
		QueueName    string `yaml:"queue_name"`
		ExchangeName string `yaml:"exchange_name"`
		ExchangeKind string `yaml:"exchange_kind"`
		Declare      bool
	}
}

var Conf *Config
var configFile string

func init() {
	flag.StringVar(&configFile, "config-file", "/etc/microservice-email.yml", "Configuration file path")
}

func ReadConfig() {
	fileData, err := ioutil.ReadFile(configFile)
	utils.CheckException(err)

	Conf = &Config{}

	err = yaml.Unmarshal(fileData, Conf)
	utils.CheckException(err)
}
