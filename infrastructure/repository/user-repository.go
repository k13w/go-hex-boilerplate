package repository

import (
	"context"
	"database/sql"
	"github.com/k13w/go-boilerplate/core/model"
)

type UserRepositoryImpl struct {
	Db *sql.DB
}

type IUserRepository interface {
	Save(ctx context.Context, user model.User)
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (b UserRepositoryImpl) Save(ctx context.Context, user model.User) {
	//tx, err := b.Db.Begin()
	//helper2.PanicIfError(err)
	//defer helper2.CommitOrRollback(tx)
	//
	//createUserSql := `INSERT INTO "user" (nome, id) VALUES ($1, DEFAULT)`
	//
	//_, err = tx.ExecContext(ctx, createUserSql, user.Name)
	//helper2.PanicIfError(err)
}
