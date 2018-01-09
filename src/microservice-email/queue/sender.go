package queue

import (
	"github.com/savsgio/go-logger"
	"github.com/streadway/amqp"
)

var msgContentType = "text/plain"

func Send(host string, queueName string, exchangeName string, exchangeKind string, msg []byte) error {
	conn, ch := New(host, queueName, exchangeName, exchangeKind, false)
	defer conn.Close()
	defer ch.Close()

	err := ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  msgContentType,
			Body:         msg,
		})

	logger.Debugf("Sent message: %s", string(msg))

	return err
}
