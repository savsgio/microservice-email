package v1

import (
	"encoding/json"
	"errors"
	"microservice-email/lib"
	"microservice-email/utils"
	"strings"

	"github.com/savsgio/go-logger"
	"github.com/valyala/fasthttp"
)

type customRequestHandler func(ctx *fasthttp.RequestCtx) error
type middleware func(ctx *fasthttp.RequestCtx) (int, error)

var middlewareList = []middleware{
	checkParamsMiddleware,
}

func checkParamsMiddleware(ctx *fasthttp.RequestCtx) (int, error) {
	params := &lib.Email{}

	if err := json.Unmarshal(ctx.PostBody(), params); err != nil {
		return fasthttp.StatusBadRequest, err
	}

	if len(params.To) == 0 {
		return fasthttp.StatusBadRequest, errors.New("invalid 'to' value")
	}

	for _, to := range params.To {
		if !strings.Contains(to, "@") {
			return fasthttp.StatusBadRequest, errors.New("invalid email " + to)
		}
	}

	if len(params.Subject) == 0 {
		return fasthttp.StatusBadRequest, errors.New("invalid 'subject' value")
	} else if params.ContentType != "plain/text" && params.ContentType != "text/html" {
		return fasthttp.StatusBadRequest, errors.New("invalid 'content_type' value")
	}

	return fasthttp.StatusOK, nil
}

func Middleware(next customRequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		logger.Debugf("%s %s", ctx.Method(), ctx.URI())

		for _, middleware := range middlewareList {
			if statusCode, err := middleware(ctx); err != nil {
				logger.Errorf("Msg: %v | RequestUri: %s", err, ctx.URI().String())

				ctx.SetStatusCode(statusCode)
				ctx.SetContentType("application/json")
				ctx.ResetBody()
				json.NewEncoder(ctx).Encode(utils.Json{"Error": err.Error()})

				return
			}
		}

		if err := next(ctx); err != nil {
			logger.Error(err)
		}
	})
}
