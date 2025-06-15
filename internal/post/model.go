package post

import "time"

type MainPagePost struct {
	ID       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	PostImg  string `db:"post_img" json:"post_img"`
	Username string `db:"username" json:"username"`
}

type PostWithComments struct {
	ID             int        `db:"post_id" json:"id"`
	Title          string     `db:"title" json:"title"`
	PostImg        string     `db:"post_img" json:"post_img"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	PostAuthor     string     `db:"post_author" json:"post_author"`
	CommentText    *string    `db:"comment_text" json:"comment_text,omitempty"`
	CommentCreated *time.Time `db:"comment_created_at" json:"comment_created_at,omitempty"`
	CommentUserID  *int       `db:"comment_user_id" json:"comment_user_id,omitempty"`
	CommentAuthor  *string    `db:"comment_author" json:"comment_author,omitempty"`
}
