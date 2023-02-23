package validations

import "github.com/go-playground/validator"

type Validate struct {
	*validator.Validate
}

func NewVaLidate() *validator.Validate {
	return validator.New()
}
