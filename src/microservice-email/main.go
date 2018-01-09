package main

import (
	"flag"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/savsgio/go-logger"
	"github.com/valyala/fasthttp"
	"microservice-email/api"
	"microservice-email/lib"
	"microservice-email/queue"
	"os"
)

var PORT = os.Getenv("PORT")

func init() {
	logLevel := flag.String("log-level", logger.WARNING, "Log level")

	// Parse only in main.go
	flag.Parse()
	// ---------------------

	logger.Setup(*logLevel)
	lib.ReadConfig()
}

func startApi() {
	router := fasthttprouter.New()
	router.POST("/", api.Index)

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
	go queue.StartConsumer(
		rabbitmqConf.Host,
		rabbitmqConf.QueueName,
		rabbitmqConf.ExchangeName,
		rabbitmqConf.ExchangeKind,
		rabbitmqConf.Declare)

	// Web API
	startApi()
}
