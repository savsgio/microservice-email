package main

import (
	"flag"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/savsgio/go-logger"
	"github.com/valyala/fasthttp"
	"microservice-email/api"
	"microservice-email/lib"
	"os"
)

var PORT = os.Getenv("PORT")

func init() {
	var logLevel string

	flag.StringVar(&logLevel, "log-level", logger.WARNING, "Log level")
	flag.StringVar(&lib.ConfigFilePath, "config-file", "/etc/microservice-email.yml", "Configuration file path")
	flag.Parse()

	logger.Setup(logLevel)
	lib.ReadConfig()
}

func StartApi() {
	router := fasthttprouter.New()
	router.POST("/api/v1/", api.V1)

	server := &fasthttp.Server{
		Name:    "MicroService Email",
		Handler: router.Handler,
	}

	logger.Debugf("Listening in http://localhost:%s...", PORT)
	logger.Fatal(server.ListenAndServe(fmt.Sprintf(":%s", PORT)))
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
	StartApi()
}
