package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/savsgio/go-logger/v2"
	"github.com/savsgio/microservice-email/internal/api"
	"github.com/savsgio/microservice-email/internal/lib"
)

var version, build string

func init() {
	var logLevel string
	var showVersion bool

	flag.StringVar(&logLevel, "log-level", logger.INFO, "Log level")
	flag.StringVar(&lib.ConfigFilePath, "config-file", "/etc/microservice-email.conf.yml", "Configuration file path")
	flag.BoolVar(&showVersion, "version", false, "Print version of service")
	flag.Parse()

	if showVersion {
		fmt.Println("Microservice-email:")
		fmt.Printf("  Version: %s\n", version)
		fmt.Printf("  Build: %s\n\n", build)
		fmt.Printf("Go version: %s\n", runtime.Version())
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
