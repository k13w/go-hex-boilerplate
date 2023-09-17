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
	createUserSql := `INSERT INTO public.users (name, class) VALUES ($1, $2)`

	err := database.NewStatement(ctx, createUserSql, user.Name, user.Class).Execute()
	if err != nil {
		return
	}
}
