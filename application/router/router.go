package router

import (
	"github.com/julienschmidt/httprouter"
	controller "github.com/k13w/go-boilerplate/application/controllers"
)

func NewRouter(userController *controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", userController.Create)

	return router
}
