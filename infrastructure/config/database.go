package database

import (
	"context"
	"database/sql"
	"github.com/k13w/go-boilerplate/core/helper"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5600
	user     = "postgres"
	password = "postgres"
	dbName   = "postgres"
)

var instance *sql.DB

func InitDBConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://go_boilerplate_database_user:V0jqW3FJseDAzhJLyKJUKIKZ5jB2VjMa@dpg-ck3gssnqj8ts73cdmnc0-a:5432/go_boilerplate_database?sslmode=disable")
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	println("Connected to database!!")

	instance = db
	return db
}

// Statement struct
type Statement struct {
	ctx   context.Context
	query string
	args  []interface{}
}

// NewStatement create a new pointer to Statement struct
func NewStatement(ctx context.Context, query string, params ...interface{}) *Statement {
	return &Statement{ctx, query, params}
}

// Execute apply statement in database
func (s *Statement) Execute() error {
	tx, err := instance.Begin()
	helper.PanicIfError(err)
	stmt, err := s.createStatement()
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(s.args...); err != nil {
		errRollback := tx.Rollback()
		helper.PanicIfError(errRollback)
		panic(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *Statement) createStatement() (*sql.Stmt, error) {
	return instance.PrepareContext(s.ctx, s.query)
}
