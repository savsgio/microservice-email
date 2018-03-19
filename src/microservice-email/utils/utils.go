package utils

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type Json map[string]interface{}

func CheckException(err error) {
	if err != nil {
		panic(err)
	}
}

func JsonResponse(ctx *fasthttp.RequestCtx, response Json, statusCode ...int) error {
	ctx.SetContentType("application/json")

	if len(statusCode) > 0 {
		ctx.SetStatusCode(statusCode[0])
	}

	ctx.ResetBody()
	return json.NewEncoder(ctx).Encode(response)
}
