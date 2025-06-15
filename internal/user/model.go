package user

import "glass-photo/internal/post"

type UserPage struct {
	Id           int64             `db:"id" json:"id"`
	Username     string            `db:"username" json:"username"`
	ProfilePhoto string            `db:"profile" json:"profile"`
	Posts        post.MainPagePost `db:"posts" json:"posts"`
}
