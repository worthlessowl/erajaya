package db

import (
	"database/sql"
)

type databaseInterface interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
}

type Database struct {
	db databaseInterface
}

func New(db databaseInterface) Database {
	return Database{db: db}
}
