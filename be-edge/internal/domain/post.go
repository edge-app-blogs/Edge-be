package domain

import "time"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Creator   string    `json:"creator"`
	Category  string    `json:"category"`
}
