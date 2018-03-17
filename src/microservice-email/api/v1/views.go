package v1

import (
	"encoding/json"
	"errors"
	"microservice-email/lib"
	"microservice-email/utils"
	"strings"

	"github.com/valyala/fasthttp"
)

func validEmailParams(m *lib.Email) (string, bool) {
	if len(m.To) == 0 || !strings.Contains(m.To, "@") {
		return "Invalid 'to' value...", true
	} else if len(m.Subject) == 0 {
		return "Invalid 'subject' value...", true
	} else if m.ContentType != "plain/text" && m.ContentType != "text/html" {
		return "Invalid email ContentType...", true
	}

	return "", false
}

// V1 is a view that receive a request and send an email
func V1(ctx *fasthttp.RequestCtx) error {
	ctx.SetContentType("application/json")

	body := ctx.PostBody()
	params := &lib.Email{}

	// Check if valid json body
	if err := json.Unmarshal(body, params); err != nil {
		return err
	}

	if errorMsg, valid := validEmailParams(params); valid {
		return errors.New(errorMsg)
	}

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

	if err := rmq.Send(body); err != nil {
		return err
	}

	response := utils.Json{"Status": "OK"}

	return json.NewEncoder(ctx).Encode(response)

}
