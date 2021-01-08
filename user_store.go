package main

import (
	"database/sql"
	"time"
)

type user struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (u *user) findByID(db *sql.DB) error {
	return db.QueryRow("SELECT id, name FROM users WHERE id = ?", u.ID).Scan(&u.ID, &u.Name)
}

func findAll(db *sql.DB) ([]user, error) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (u *user) store(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO users (name) VALUES (?)", u.Name)
	return err
}

func (u *user) update(db *sql.DB) error {
	_, err := db.Exec("UPDATE users SET name = ? WHERE id = ?", u.Name, u.ID)
	return err
}

func (u *user) delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", u.ID)
	return err
}
