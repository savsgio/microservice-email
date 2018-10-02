package api

import (
	"microservice-email/lib"

	"github.com/savsgio/atreugo"
)

// SendEmailView is a view that receive a request and send an email
func sendEmailView(ctx *atreugo.RequestCtx) error {
	rabbitmqConf := lib.Conf.RabbitMQ
	rmq, err := lib.NewRabbitMQ(
		rabbitmqConf.Host,
		rabbitmqConf.User,
		rabbitmqConf.Password,
		rabbitmqConf.QueueName,
		rabbitmqConf.ExchangeName,
		rabbitmqConf.ExchangeKind,
		false,
	)
	if err != nil {
		return err
	}

	body := ctx.PostBody()
	if err := rmq.Send(body); err != nil {
		return err
	}

	return ctx.JSONResponse(atreugo.JSON{"status": "ok"})
}
