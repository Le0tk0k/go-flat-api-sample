package main

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type app struct {
	router *mux.Router
	db     *sqlx.DB
}

func (a *app) init() {

}

func (a *app) initRoutes() {

}

func (a *app) run() {

}

func newDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/sample?charset=utf8")
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	return db, nil
}
