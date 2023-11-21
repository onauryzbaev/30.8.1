package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// Post - публикация.
type Post struct {
	ID          int
	Title       string
	Content     string
	AuthorID    int
	AuthorName  string
	CreatedAt   int64
	PublishedAt int64
}

// DB - структура, представляющая соединение с базой данных.
type DB struct {
	conn *sql.DB
}

// NewDB создает новый экземпляр DB.
func NewDB(db *sql.DB) *DB {
	return &DB{conn: db}
}

// Posts возвращает все публикации.
func (db *DB) Posts() ([]Post, error) {
	rows, err := db.conn.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorName,
			&post.CreatedAt,
			&post.PublishedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// AddPost добавляет новую публикацию.
func (db *DB) AddPost(post Post) error {
	_, err := db.conn.Exec("INSERT INTO posts (title, content, author_id, author_name, created_at, published_at) VALUES (?, ?, ?, ?, ?, ?)",
		post.Title, post.Content, post.AuthorID, post.AuthorName, post.CreatedAt, post.PublishedAt)
	return err
}

// UpdatePost обновляет существующую публикацию.
func (db *DB) UpdatePost(post Post) error {
	result, err := db.conn.Exec("UPDATE posts SET title=?, content=?, author_id=?, author_name=?, created_at=?, published_at=? WHERE id=?",
		post.Title, post.Content, post.AuthorID, post.AuthorName, post.CreatedAt, post.PublishedAt, post.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("post not found")
	}

	return nil
}

// DeletePost удаляет публикацию по ID.
func (db *DB) DeletePost(id int) error {
	result, err := db.conn.Exec("DELETE FROM posts WHERE id=?", id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("post not found")
	}

	return nil
}
