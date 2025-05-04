package main

import (
	"books_service/controller"
	"books_service/middleware"
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
	middleware := middleware.NewJwtMiddleware(handler)
	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: middleware,
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
