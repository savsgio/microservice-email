package queue

import (
	"fmt"
	"github.com/savsgio/go-logger"
	"github.com/streadway/amqp"
	"microservice-email/utils"
)

func New(host string, queueName string, exchangeName string, exchangeKind string, declare bool) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s/", host))
	utils.CheckException(err)

	ch, err := conn.Channel()
	utils.CheckException(err)

	if declare {
		logger.Debugf("Declaring exchange: %s", exchangeName)
		err = ch.ExchangeDeclare(
			exchangeName,
			exchangeKind,
			true,
			false,
			false,
			false,
			nil,
		)
		utils.CheckException(err)

		logger.Debugf("Declaring queue: %s", queueName)
		_, err = ch.QueueDeclare(
			queueName,
			true,
			false,
			false,
			false,
			nil,
		)
		utils.CheckException(err)

		logger.Debug("Setting RabbitMQ channel Qos...")
		err = ch.Qos(
			1,
			0,
			false,
		)
		utils.CheckException(err)
	} else {
		logger.Debugf("Binding queue: %s", queueName)
		err := ch.QueueBind(
			queueName,
			"",
			exchangeName,
			false,
			nil,
		)
		utils.CheckException(err)
	}

	return conn, ch
}
