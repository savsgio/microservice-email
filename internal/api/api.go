package api

import (
	"strconv"

	"github.com/savsgio/atreugo/v10"
)

type Api struct {
	server *atreugo.Atreugo
}

func New(port int) *Api {
	if port == 0 {
		port = 8000 // Default port
	}

	api := &Api{
		server: atreugo.New(atreugo.Config{
			Addr:             "0.0.0.0" + strconv.Itoa(port),
			GracefulShutdown: true,
		}),
	}

	api.setRoutes()
	api.registerMiddlewares()

	return api
}

func (api *Api) setRoutes() {
	api.server.Path("POST", "/api/v1/", sendEmailView)
}

func (api *Api) registerMiddlewares() {
	api.server.UseBefore(checkParamsMiddleware)
}

func (api *Api) ListenAndServe() error {
	return api.server.ListenAndServe()
}
