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
	repository.NewBookRepositoryImpl,
	service.NewBookServiceImpl,
	controller.NewBookController,
	NewRouter,
	NewServer, wire.Bind(new(http.Handler), new(*httprouter.Router)),
)

func InitServer() (*http.Server, func(), error) {
	wire.Build(ServerSet)
	return nil, nil, nil
}
