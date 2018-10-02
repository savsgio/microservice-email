package api

import (
	"encoding/json"
	"errors"
	"microservice-email/lib"
	"strings"

	"github.com/savsgio/atreugo"
	"github.com/valyala/fasthttp"
)

func checkParamsMiddleware(ctx *atreugo.RequestCtx) (int, error) {
	params := new(lib.Email)

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
