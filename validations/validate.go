package validations

import "github.com/go-playground/validator/v10"

type Validate struct {
	*validator.Validate
}

func NewVaLidate() *validator.Validate {
	return validator.New()
}
