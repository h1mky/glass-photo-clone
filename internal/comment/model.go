package comment

import "time"

type Comments struct {
	IdComment int       `json:"id" db:"id"`
	UserId    int       `json:"userId" db:"user_id"`
	UserImg   string    `json:"userImg" db:"user_img"`
	PostId    int       `json:"postId" db:"post_id"`
	UserName  string    `json:"userName" db:"username"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CommentPost struct {
	Content string `json:"content" db:"content" validate:"required"`
}
type CommentUser struct {
	IdComment int       `json:"id" db:"id"`
	UserId    int       `json:"userId" db:"user_id"`
	PostId    int       `json:"postId" db:"post_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
