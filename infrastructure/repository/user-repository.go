package repository

import (
	"context"
	"github.com/k13w/go-boilerplate/core/model"
	database "github.com/k13w/go-boilerplate/infrastructure/config"
)

type UserRepositoryImpl struct{}

type IUserRepository interface {
	Save(ctx context.Context, user model.User)
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (b UserRepositoryImpl) Save(ctx context.Context, user model.User) {
	createUserSql := `INSERT INTO public.users (name) VALUES ($1)`

	err := database.NewStatement(ctx, createUserSql, user.Name).Execute()
	if err != nil {
		return
	}
}
