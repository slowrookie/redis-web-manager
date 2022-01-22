package api

import (
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed initialize.sql
var initializeSQL string

var (
	DB *sql.DB
)

// InitializeDB initialize database
func InitializeDB(dsn string) error {
	if nil != DB {
		return nil
	}

	var err error
	DB, err = sql.Open("sqlite3", dsn)
	if nil != err {
		return err
	}
	_, err = DB.Exec(initializeSQL)
	if nil != err {
		return err
	}
	return err
}
