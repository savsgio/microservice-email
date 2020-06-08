package lib

import (
	"encoding/json"
	"fmt"

	"github.com/savsgio/go-logger/v2"
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

func NewRabbitMQ(host string, user string, password string, queueName string, exchangeName string, exchangeKind string, declare bool) (*RabbitMQ, error) {
	var err error
	rmq := &RabbitMQ{Host: host, QueueName: queueName, ExchangeName: exchangeName, ExchangeKind: exchangeKind, Declare: declare}

	rmq.Connection, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", user, password, host))
	if err != nil {
		panic(err)
	}

	rmq.Channel, err = rmq.Connection.Channel()
	if err != nil {
		panic(err)
	}

	if declare {
		err = rmq.exchangeAndQueueDeclare()
	} else {
		err = rmq.queueBind()
	}

	return rmq, err
}

func (rmq *RabbitMQ) exchangeAndQueueDeclare() error {
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
	if err != nil {
		return err
	}

	logger.Debugf("Declaring queue: %s", rmq.QueueName)
	_, err = rmq.Channel.QueueDeclare(
		rmq.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	logger.Debug("Setting RabbitMQ channel Qos...")
	err = rmq.Channel.Qos(
		1,
		0,
		false,
	)

	return err
}

func (rmq *RabbitMQ) queueBind() error {
	logger.Debugf("Binding queue: %s", rmq.QueueName)
	return rmq.Channel.QueueBind(
		rmq.QueueName,
		"",
		rmq.ExchangeName,
		false,
		nil,
	)
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

	email := new(Email)
	err := json.Unmarshal(d.Body, email)
	if err != nil {
		logger.Error(err)
	}

	err = email.Send()
	if err != nil {
		logger.Error(err)
	} else {
		logger.Debug("Email send successfully...")
	}

	d.Ack(false)
}

func (rmq *RabbitMQ) StartConsumer() error {
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
	if err != nil {
		return err
	}

	logger.Info("[*] Waiting for messages. To exit press CTRL+C")
	for d := range msgs {
		callback(d)
	}

	return nil
}
