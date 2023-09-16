package main

import (
	controller "github.com/k13w/go-boilerplate/application/controllers"
	"github.com/k13w/go-boilerplate/application/router"
	"github.com/k13w/go-boilerplate/core/helper"
	"github.com/k13w/go-boilerplate/infrastructure/config"
	"net/http"
)

func init() {
	database.InitDBConnection()
}

func main() {
	userController := controller.NewUserController()

	routes := router.NewRouter(userController)

	server := http.Server{Addr: "localhost:7000", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
