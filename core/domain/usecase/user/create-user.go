package usecase

import (
	"context"
	"github.com/k13w/go-boilerplate/core/model"
	"github.com/k13w/go-boilerplate/core/utils/request"
	"github.com/k13w/go-boilerplate/infrastructure/repository"
)

type ICreateUserusecase interface {
	Execute(ctx context.Context, request request.UserCreateRequest)
}

type CreateUserUsecase struct {
	UserRepository repository.IUserRepository
}

func NewUserUseCase() *CreateUserUsecase {
	return &CreateUserUsecase{
		UserRepository: repository.NewUserRepository(),
	}
}

func (u *CreateUserUsecase) Execute(ctx context.Context, request request.UserCreateRequest) {
	user := model.User{
		Name: request.Name,
	}

	u.UserRepository.Save(ctx, user)
}
