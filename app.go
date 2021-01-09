package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type app struct {
	router *mux.Router
	db     *sql.DB
}

func (a *app) init() {
	db, err := newDb()
	if err != nil {
		log.Fatal(err)
	}

	a.db = db
	a.router = mux.NewRouter()

	a.initRoutes()
}

func (a *app) initRoutes() {
	a.router.HandleFunc("/users/{id}", a.getUser).Methods("GET")
	a.router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.router.HandleFunc("/users", a.createUser).Methods("POST")
	a.router.HandleFunc("/users/{id}", a.updateUser).Methods("PUT")
	a.router.HandleFunc("/users/{id}", a.deleteUser).Methods("DELETE")

	a.router.HandleFunc("/posts/{id}", a.getPost).Methods("GET")
	a.router.HandleFunc("/posts", a.getPosts).Methods("GET")
	a.router.HandleFunc("/posts", a.createPost).Methods("POST")
	a.router.HandleFunc("/posts/{id}", a.updatePost).Methods("PUT")
	a.router.HandleFunc("/posts/{id}", a.deletePost).Methods("DELETE")
}

func (a *app) run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.router))
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
