package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

func New(msg string) error {
	return errors.New(msg)
}

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

func Cause(err error) error {
	return errors.Cause(err)
}

type MultiError struct {
	errors []error
}

func NewMultiError() *MultiError {
	return &MultiError{
		errors: make([]error, 0),
	}
}

func (me *MultiError) Append(e error) {
	me.errors = append(me.errors, e)
}

func (me *MultiError) HasErrors() bool {
	return len(me.errors) > 0
}

func (me *MultiError) Error() string {
	return fmt.Sprintf(
		"multiple errors occured: %v",
		me.errors,
	)
}
