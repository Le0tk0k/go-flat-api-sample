package main

import "database/sql"

func newDb() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/sample?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return db
}
