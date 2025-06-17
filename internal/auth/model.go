package auth

import _ "github.com/go-playground/validator/v10"

type UserInputRegister struct {
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,min=8"`
	Username string `json:"username" db:"username" validate:"required"`
}

type UserInputSignIn struct {
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required"`
}
