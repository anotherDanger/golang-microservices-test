package main

import (
	"books_service/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(ctrl *controller.BookControllerImpl) *httprouter.Router {
	r := httprouter.New()
	r.POST("/v1/book", ctrl.Create)
	r.PUT("/v1/book/:id", ctrl.Update)
	r.DELETE("/v1/book/:id", ctrl.Delete)
	r.GET("/v1/book/:id", ctrl.FindById)
	r.GET("/v1/book", ctrl.FindAll)
	return r
}

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}
}

func main() {
	server, cleanup, err := InitServer()
	if err != nil {
		panic(err)
	}

	defer cleanup()

	server.ListenAndServe()
}
