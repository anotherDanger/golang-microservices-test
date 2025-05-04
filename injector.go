//go:build wireinject
// +build wireinject

package main

import (
	"books_service/controller"
	"books_service/helper"
	"books_service/repository"
	"books_service/service"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var ServerSet = wire.NewSet(
	helper.NewDb,
	repository.NewBookRepositoryImpl, wire.Bind(new(repository.BookRepository), new(*repository.BookRepositoryImpl)),
	service.NewBookServiceImpl, wire.Bind(new(service.BookService), new(*service.BookServiceImpl)),
	controller.NewBookController, wire.Bind(new(controller.BookController), new(*controller.BookControllerImpl)),
	NewRouter,
	NewServer, wire.Bind(new(http.Handler), new(*httprouter.Router)),
)

func InitServer() (*http.Server, func(), error) {
	wire.Build(ServerSet)
	return nil, nil, nil
}
