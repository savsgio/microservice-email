package api

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/savsgio/atreugo/v11"
	"github.com/savsgio/microservice-email/internal/lib"
	"github.com/valyala/fasthttp"
)

func checkParamsMiddleware(ctx *atreugo.RequestCtx) error {
	params := new(lib.Email)

	if err := json.Unmarshal(ctx.PostBody(), params); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return err
	}

	if len(params.To) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return errors.New("invalid 'to' value")
	}

	for _, to := range params.To {
		if !strings.Contains(to, "@") {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)

			return errors.New("invalid email " + to)
		}
	}

	if len(params.Subject) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return errors.New("invalid 'subject' value")
	} else if params.ContentType != "plain/text" && params.ContentType != "text/html" {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return errors.New("invalid 'content_type' value")
	}

	return ctx.Next()
}
