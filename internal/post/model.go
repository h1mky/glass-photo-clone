package post

import "time"

type MainPagePost struct {
	ID       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	PostImg  string `db:"post_img" json:"post_img"`
	Username string `db:"username" json:"username"`
}

type PostInfo struct {
	PostID       int       `db:"post_id"`
	Title        string    `db:"title"`
	PostImg      string    `db:"post_img"`
	CreatedAt    time.Time `db:"created_at"`
	PostAuthor   string    `db:"post_author"`
	PostAuthorID int       `db:"post_author_id"`
}
