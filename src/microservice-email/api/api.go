package api

import (
	"github.com/savsgio/atreugo"
)

type Api struct {
	Server *atreugo.Atreugo
}

func New(port int) *Api {
	if port == 0 {
		port = 8000 // Default port
	}

	api := &Api{
		Server: atreugo.New(&atreugo.Config{
			Host:             "0.0.0.0",
			Port:             port,
			GracefulShutdown: true,
		}),
	}

	api.setRoutes()
	api.registerMiddlewares()

	return api
}

func (api *Api) setRoutes() {
	api.Server.Path("POST", "/api/v1/", sendEmailView)
}

func (api *Api) registerMiddlewares() {
	api.Server.UseMiddleware(checkParamsMiddleware)
}

func (api *Api) ListenAndServe() error {
	return api.Server.ListenAndServe()
}
