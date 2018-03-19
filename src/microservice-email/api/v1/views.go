package v1

import (
	"microservice-email/lib"
	"microservice-email/utils"

	"github.com/valyala/fasthttp"
)

// SendEmailView is a view that receive a request and send an email
func SendEmailView(ctx *fasthttp.RequestCtx) error {
	rabbitmqConf := lib.Conf.RabbitMQ
	rmq := lib.NewRabbitMQ(
		rabbitmqConf.Host,
		rabbitmqConf.User,
		rabbitmqConf.Password,
		rabbitmqConf.QueueName,
		rabbitmqConf.ExchangeName,
		rabbitmqConf.ExchangeKind,
		false,
	)

	body := ctx.PostBody()
	if err := rmq.Send(body); err != nil {
		return err
	}

	return utils.JsonResponse(ctx, utils.Json{"Status": "OK"})
}
