package api

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"microservice-email/lib"
)

const HttpErrorMsg = "{\"Error\": \"%v\"}"
const HttpSuccessMsg = "{\"Status\": \"OK\"}"

func validEmailParams(m *lib.Email) (string, bool) {
	if len(m.To) == 0 {
		return "Invalid 'to' value...", true
	} else if len(m.Subject) == 0 {
		return "Invalid 'subject' value...", true
	} else if m.ContentType != "plain/text" && m.ContentType != "text/html" {
		return "Invalid email ContentType...", true
	}

	return "", false
}

func V1(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	body := ctx.PostBody()
	params := &lib.Email{}

	// Check if valid json body
	err := json.Unmarshal(body, params)
	if err != nil {
		fmt.Fprintf(ctx, HttpErrorMsg, err)
		ctx.SetStatusCode(500)
		return
	}
	if errorMsg, valid := validEmailParams(params); valid {
		fmt.Fprintf(ctx, HttpErrorMsg, errorMsg)
		ctx.SetStatusCode(400)
		return
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
	err = rmq.Send(body)

	if err != nil {
		fmt.Fprintf(ctx, HttpErrorMsg, err)
		ctx.SetStatusCode(500)
	} else {
		fmt.Fprint(ctx, HttpSuccessMsg)
	}
}
