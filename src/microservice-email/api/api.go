package api

import (
	"fmt"
	"microservice-email/api/v1"

	"github.com/buaazp/fasthttprouter"
	"github.com/savsgio/go-logger"
	"github.com/valyala/fasthttp"
)

func StartApi(port string) {
	router := fasthttprouter.New()
	router.POST("/api/v1/", v1.Middleware(v1.V1))

	server := &fasthttp.Server{
		Name:    "MicroService Email",
		Handler: router.Handler,
	}

	logger.Debugf("Listening in http://localhost:%s...", port)
	logger.Fatal(server.ListenAndServe(fmt.Sprintf(":%s", port)))
}
