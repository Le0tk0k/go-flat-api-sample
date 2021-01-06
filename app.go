package main

import (
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
