package main

import (
	"database/sql"
	"fmt"
)

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
