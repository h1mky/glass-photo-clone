package comment

import "time"

type Comments struct {
	IdСomment int       `json:"id" db:"id"`
	UserId    int       `json:"userId" db:"user_id"`
	PostId    int       `json:"postId" db:"post_id"`
	UserName  string    `json:"userName" db:"user_name"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
