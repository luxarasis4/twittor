package dtos

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(request interface{}) error {
	return validator.New().Struct(request)
}
