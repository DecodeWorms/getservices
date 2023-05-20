package validations

import (
	"fmt"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

type Validate struct {
	*validator.Validate
}

func NewVaLidate() *validator.Validate {
	return validator.New()
}

//func ValidatedData(v Validate, data interface{}) []error {
//errDetails := make([]error, 0)

//err := v.Struct(data)
//if err != nil {
//for _, err := range err.(validator.ValidationErrors) {
//e := fmt.Errorf(fmt.Sprintf("user_data object: a valid %v of type %v is required, but recieved '%v' ", strings.ToLower(err.Field()), err.Kind(), err.Value()))
//errDetails = append(errDetails, e)
//}
//return errDetails
//}

//return errDetails
//}

func ValidatedData(v *validator.Validate, data interface{}) []error {
	errDetails := make([]error, 0)

	err := v.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			e := fmt.Errorf(fmt.Sprintf("user_data object: a valid %v of type %v is required, but recieved '%v' ", strings.ToLower(err.Field()), err.Kind(), err.Value()))
			errDetails = append(errDetails, e)
		}
		return errDetails
	}

	return errDetails
}
