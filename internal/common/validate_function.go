package common

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(user interface{}) error {
	err := validate.Struct(user)
	if err != nil {
		return err
	}

	return nil
}
