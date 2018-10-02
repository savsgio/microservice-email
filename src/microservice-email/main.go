package main

import (
	"flag"
	"microservice-email/api"
	"microservice-email/lib"
	"os"
	"strconv"

	"github.com/savsgio/go-logger"
)

const version = "1.2.0"

func init() {
	var logLevel string
	var showVersion bool

	flag.StringVar(&logLevel, "log-level", logger.INFO, "Log level")
	flag.StringVar(&lib.ConfigFilePath, "config-file", "/etc/microservice-email.yml", "Configuration file path")
	flag.BoolVar(&showVersion, "version", false, "Print version of service")
	flag.Parse()

	if showVersion {
		println("Version: " + version)
		os.Exit(0)
	}

	logger.SetLevel(logLevel)
	lib.ReadConfig()
}

func main() {
	rmqConf := lib.Conf.RabbitMQ
	rmq, err := lib.NewRabbitMQ(
		rmqConf.Host,
		rmqConf.User,
		rmqConf.Password,
		rmqConf.QueueName,
		rmqConf.ExchangeName,
		rmqConf.ExchangeKind,
		rmqConf.Declare,
	)
	if err != nil {
		panic(err)
	}

	go rmq.StartConsumer()

	// Web API
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	err = api.New(port).ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
