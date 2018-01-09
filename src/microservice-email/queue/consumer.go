package queue

import (
	"encoding/json"
	"github.com/savsgio/go-logger"
	"github.com/streadway/amqp"
	"microservice-email/lib"
	"microservice-email/utils"
)

func callback(d amqp.Delivery) {
	logger.Debugf("Received a message: %s", d.Body)

	email := &lib.Email{}
	err := json.Unmarshal(d.Body, email)
	utils.CheckException(err)

	err = email.Send()
	utils.CheckException(err)

	logger.Debug("Email send successfully...")

	d.Ack(false)
}

func StartConsumer(host string, queueName string, exchangeName string, exchangeKind string, declare bool) {
	conn, ch := New(host, queueName, exchangeName, exchangeKind, declare)
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.CheckException(err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			callback(d)
		}
	}()

	logger.Info("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
