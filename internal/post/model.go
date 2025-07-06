package post

import (
	"database/sql"
	"time"
)

type MainPagePost struct {
	ID       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	PostImg  string `db:"post_img" json:"post_img"`
	Username string `db:"username" json:"username"`
}

type PostInfo struct {
	PostID          int            `db:"post_id" json:"post_id"`
	Title           string         `db:"title" json:"title"`
	PostImg         string         `db:"post_img" json:"post_img"`
	CreatedAt       time.Time      `db:"created_at" json:"created_at"`
	PostAuthor      string         `db:"post_author" json:"post_author"`
	PostAuthorID    int            `db:"post_author_id" json:"post_author_id"`
	PostAuthorIMG   string         `db:"post_author_img" json:"post_author_img"`
	PostDescription sql.NullString `db:"description" json:"description"`
}

type CreatePostRequest struct {
	Title       string `json:"title" db:"title" validate:"required"`
	PostImg     string `json:"post_img" db:"post_img" validate:"required,url"`
	Description string `json:"description" db:"description" validate:"omitempty,max=64"`
}
