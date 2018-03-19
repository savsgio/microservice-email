package main

import (
	"flag"
	"microservice-email/api"
	"microservice-email/lib"
	"os"

	"github.com/savsgio/go-logger"
)

func init() {
	var logLevel string

	flag.StringVar(&logLevel, "log-level", logger.WARNING, "Log level")
	flag.StringVar(&lib.ConfigFilePath, "config-file", "/etc/microservice-email.yml", "Configuration file path")
	flag.Parse()

	logger.Setup(logLevel)
	lib.ReadConfig()
}

func main() {
	// RabbitMQ Consumer
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
