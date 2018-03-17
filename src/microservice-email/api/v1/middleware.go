package v1

import (
	"github.com/savsgio/go-logger"
	"github.com/valyala/fasthttp"
)

type customRequestHandler func(ctx *fasthttp.RequestCtx) error
type middleware func(ctx *fasthttp.RequestCtx) (int, error)

var middlewareList = []middleware{
	// authMiddleware
}

// func authMiddleware(ctx *fasthttp.RequestCtx) (int, error) {
// 	// Example
// 	if strings.Contains(ctx.URI().String(), "error") {
// 		return fasthttp.StatusBadRequest, errors.New("invalid request")
// 	}

// 	return 0, nil
// }

func Middleware(next customRequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		logger.Debugf("%s %s", ctx.Method(), ctx.URI())

		for _, middleware := range middlewareList {
			if statusCode, err := middleware(ctx); err != nil {
				logger.Errorf("Msg: %v | RequestUri: %s", err, ctx.URI().String())

				ctx.SetStatusCode(statusCode)
				ctx.ResetBody()
				ctx.Write([]byte(err.Error()))

				return
			}
		}

		if err := next(ctx); err != nil {
			logger.Error(err)
		}
	})
}
