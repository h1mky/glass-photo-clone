package common

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(input any) error {
	err := validate.Struct(input)
	if err != nil {
		return err
	}

	return nil
}
