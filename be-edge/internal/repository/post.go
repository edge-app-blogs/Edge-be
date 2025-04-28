package repository

import (
	"be-edge/internal/domain"
	"database/sql"
	"fmt"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *domain.Post) error {
	result, err := r.db.Exec("INSERT INTO posts (title, content, created_at, creator, category) VALUES (?, ?, ?, ?, ?)", 
		post.Title, post.Content, post.CreatedAt, post.Creator, post.Category)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	post.ID = int(id)

	return nil
}


func (r *PostRepository) GetAll() ([]domain.Post, error) {
	rows, err := r.db.Query("SELECT id, title, content, creator, category  FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var post domain.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Creator, &post.Category); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
}

func (r *PostRepository) GetByID(id int) (*domain.Post, error) {
	row := r.db.QueryRow("SELECT id, title, content, creator, category FROM posts WHERE id = ?", id)
	var post domain.Post
	if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Creator, &post.Category); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving post: %v", err)
	}
	return &post, nil
}

func (r *PostRepository) Update(post *domain.Post) error {
	_, err := r.db.Exec(
		"UPDATE posts SET title = ?, content = ?, creator = ?, category = ?, created_at = ? WHERE id = ?",
		post.Title, post.Content, post.Creator, post.Category, post.CreatedAt, post.ID,
	)
	return err
}

