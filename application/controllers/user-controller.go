package controller

import (
	"github.com/julienschmidt/httprouter"
	usecase "github.com/k13w/go-boilerplate/core/domain/usecase/user"
	"github.com/k13w/go-boilerplate/core/helper"
	"github.com/k13w/go-boilerplate/core/utils/request"
	"github.com/k13w/go-boilerplate/core/utils/response"
	"net/http"
)

type UserController struct {
	CreateUserUsecase usecase.ICreateUserusecase
}

func NewUserController() *UserController {
	return &UserController{
		CreateUserUsecase: usecase.NewUserUseCase(),
	}
}

func (controller *UserController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	userCreateRequest := request.UserCreateRequest{}
	helper.ReadRequestBody(requests, &userCreateRequest)

	controller.CreateUserUsecase.Execute(requests.Context(), userCreateRequest)
	webResponse := response.WebResponse{
		Code:   201,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}
