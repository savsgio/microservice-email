package api

import (
	"fmt"
	"microservice-email/api/v1"

	"github.com/buaazp/fasthttprouter"
	"github.com/savsgio/go-logger"
	"github.com/valyala/fasthttp"
)

type Api struct {
	Addr   string
	Server *fasthttp.Server
	Router *fasthttprouter.Router
}

func New(port string) *Api {
	router := fasthttprouter.New()

	if len(port) == 0 {
		port = "8000" // Default port
	}

	api := &Api{
		Addr:   fmt.Sprintf("0.0.0.0:%s", port),
		Router: router,
		Server: &fasthttp.Server{
			Handler: router.Handler,
			Name:    "MicroService Email",
		},
	}

	api.setRoutesV1()

	return api
}

func (api *Api) setRoutesV1() {
	api.Router.POST("/api/v1/", v1.Middleware(v1.SendEmailView))
}

func (api *Api) ListenAndServe() {
	logger.Infof("Listening on: http://%s/", api.Addr)
	logger.Fatal(api.Server.ListenAndServe(api.Addr))
}
