package main

import (
	"database/sql"
	"time"
)

type post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (p *post) findByID(db *sql.DB) error {
	return db.QueryRow("SELECT id, title, body, author_id FROM posts WHERE id = ?", p.ID).Scan(&p.ID, &p.Title, &p.Body, &p.AuthorID)
}

func findAllPosts(db *sql.DB) ([]post, error) {
	rows, err := db.Query("SELECT id, title, body, author_id FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []post

	for rows.Next() {
		var p post
		if err := rows.Scan(&p.ID, &p.Title, &p.Body, &p.AuthorID); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (p *post) store(db *sql.DB, authorID int) error {
	_, err := db.Exec("INSERT INTO posts (title, body, author_id) VALUES (?, ?, ?)", p.Title, p.Body, authorID)
	return err
}

func (p *post) update(db *sql.DB) error {
	_, err := db.Exec("UPDATE posts SET title = ?, body = ? WHERE id = ?", p.Title, p.Body, p.ID)
	return err
}

func (p *post) delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", p.ID)
	return err
}
