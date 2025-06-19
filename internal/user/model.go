package user

import (
	"database/sql"
)

type UserPage struct {
	Id           int64          `db:"id" json:"id"`
	Username     string         `db:"username" json:"username"`
	ProfilePhoto sql.NullString `db:"user_img" json:"user_img"`
	Description  sql.NullString `db:"description" json:"description"`
}
type MainPageRegistered struct {
	Id           int            `db:"id" json:"id"`
	ProfilePhoto sql.NullString `db:"user_img" json:"user_img"`
}

type UserPatchRequest struct {
	UserImg     *string `json:"user_img,omitempty" db:"user_img"`
	Description *string `json:"description,omitempty" db:"description"`
}
