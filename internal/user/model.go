package user

import (
	"database/sql"
	"glass-photo/internal/post"
)

type UserPage struct {
	Id           int64               `db:"id" json:"id"`
	Username     string              `db:"username" json:"username"`
	ProfilePhoto sql.NullString      `db:"user_img" json:"user_img"`
	Description  sql.NullString      `db:"description" json:"description"`
	Posts        []post.MainPagePost `db:"posts" json:"posts"`
}
type MainPageRegistered struct {
	Id           int            `db:"id" json:"id"`
	ProfilePhoto sql.NullString `db:"user_img" json:"user_img"`
}
