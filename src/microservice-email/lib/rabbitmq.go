package lib

import (
	"encoding/json"
	"fmt"
	"microservice-email/utils"

	"github.com/savsgio/go-logger"
	"github.com/streadway/amqp"
)

const MsgContentType = "text/plain"

type RabbitMQ struct {
	Host         string
	QueueName    string
	ExchangeName string
	ExchangeKind string
	Declare      bool
	Connection   *amqp.Connection
	Channel      *amqp.Channel
}

func NewRabbitMQ(host string, user string, password string, queueName string, exchangeName string, exchangeKind string, declare bool) *RabbitMQ {
	var err error
	rmq := &RabbitMQ{Host: host, QueueName: queueName, ExchangeName: exchangeName, ExchangeKind: exchangeKind, Declare: declare}

	rmq.Connection, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", user, password, host))
	utils.CheckException(err)

	rmq.Channel, err = rmq.Connection.Channel()
	utils.CheckException(err)

	if declare {
		rmq.exchangeAndQueueDeclare()
	} else {
		rmq.queueBind()
	}

	return rmq
}

func (rmq *RabbitMQ) exchangeAndQueueDeclare() {
	logger.Debugf("Declaring exchange: %s", rmq.ExchangeName)
	err := rmq.Channel.ExchangeDeclare(
		rmq.ExchangeName,
		rmq.ExchangeKind,
		true,
		false,
		false,
		false,
		nil,
	)
	utils.CheckException(err)

	logger.Debugf("Declaring queue: %s", rmq.QueueName)
	_, err = rmq.Channel.QueueDeclare(
		rmq.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	utils.CheckException(err)

	logger.Debug("Setting RabbitMQ channel Qos...")
	err = rmq.Channel.Qos(
		1,
		0,
		false,
	)
	utils.CheckException(err)
}

func (rmq *RabbitMQ) queueBind() {
	logger.Debugf("Binding queue: %s", rmq.QueueName)
	err := rmq.Channel.QueueBind(
		rmq.QueueName,
		"",
		rmq.ExchangeName,
		false,
		nil,
	)
	utils.CheckException(err)
}

func (rmq *RabbitMQ) Send(msg []byte) error {
	defer rmq.Channel.Close()
	defer rmq.Connection.Close()

	err := rmq.Channel.Publish(
		rmq.ExchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  MsgContentType,
			Body:         msg,
		})

	logger.Debugf("Sent message: %s", string(msg))

	return err
}

func callback(d amqp.Delivery) {
	logger.Debugf("Received a message: %s", d.Body)

	email := &Email{}
	err := json.Unmarshal(d.Body, email)
	utils.CheckException(err)

	err = email.Send()
	if err != nil {
		logger.Error(err)
	} else {
		logger.Debug("Email send successfully...")
	}

	d.Ack(false)
}

func (rmq *RabbitMQ) StartConsumer() {
	defer rmq.Channel.Close()
	defer rmq.Connection.Close()

	msgs, err := rmq.Channel.Consume(
		rmq.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.CheckException(err)

	logger.Info("[*] Waiting for messages. To exit press CTRL+C")
	for d := range msgs {
		callback(d)
	}
}
