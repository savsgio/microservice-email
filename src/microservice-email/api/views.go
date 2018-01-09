package api

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"microservice-email/lib"
	"microservice-email/queue"
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

func Index(ctx *fasthttp.RequestCtx) {
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
	err = queue.Send(
		rabbitmqConf.Host,
		rabbitmqConf.QueueName,
		rabbitmqConf.ExchangeName,
		rabbitmqConf.ExchangeKind,
		body)

	if err != nil {
		fmt.Fprintf(ctx, HttpErrorMsg, err)
		ctx.SetStatusCode(500)

	} else {
		fmt.Fprint(ctx, HttpSuccessMsg)
	}
}
