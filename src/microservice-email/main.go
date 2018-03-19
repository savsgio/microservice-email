package main

import (
	"flag"
	"microservice-email/api"
	"microservice-email/lib"
	"os"

	"github.com/savsgio/go-logger"
)

const version = "1.1.0"

func init() {
	var logLevel string
	var showVersion bool

	flag.StringVar(&logLevel, "log-level", logger.WARNING, "Log level")
	flag.StringVar(&lib.ConfigFilePath, "config-file", "/etc/microservice-email.yml", "Configuration file path")
	flag.BoolVar(&showVersion, "version", false, "Print version of service")
	flag.Parse()

	if showVersion {
		println("Version: " + version)
		os.Exit(0)
	}

	logger.Setup(logLevel)
	lib.ReadConfig()
}

func main() {
	// Email Sender
	rabbitmqConf := lib.Conf.RabbitMQ

	go lib.NewRabbitMQ(
		rabbitmqConf.Host,
		rabbitmqConf.User,
		rabbitmqConf.Password,
		rabbitmqConf.QueueName,
		rabbitmqConf.ExchangeName,
		rabbitmqConf.ExchangeKind,
		rabbitmqConf.Declare,
	).StartConsumer()

	// Web API
	api.New(os.Getenv("PORT")).ListenAndServe()
}
