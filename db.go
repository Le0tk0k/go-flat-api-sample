package main

import (
	"database/sql"
	"fmt"
)

func newDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/sample?charset=utf8")
	if err != nil {
		return nil, fmt.Errorf("failed to Open mysql: %w", err)
	}
	return db, nil
}
