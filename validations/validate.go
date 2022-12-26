package validations

import "github.com/go-playground/validator/v10"

type Validate struct {
	val *validator.Validate
}

func NewVaLidate() *validator.Validate {
	return validator.New()
}
