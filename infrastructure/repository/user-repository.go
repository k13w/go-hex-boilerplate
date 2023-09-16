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
	//sqlArch := database.InitDBConnection()
	//println("call user repository")
	//tx, err := sqlArch.Begin()
	//helper.PanicIfError(err)
	//defer helper.CommitOrRollback(tx)
	//
	createUserSql := `INSERT INTO public.users (name) VALUES ($1)`
	//
	//_, err = tx.ExecContext(ctx, createUserSql, user.Name)
	//helper.PanicIfError(err)
	err := database.NewStatement(ctx, createUserSql, user.Name).Execute()
	if err != nil {
		return
	}
}
