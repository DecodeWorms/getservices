package hashpassword

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	MinCost                      = 4
	MaxCost                      = 14
	DefaultCost                  = 10
	ErrHashTooShort              = errors.New("crypto/bcrypt: hashedSecret too short to be a bcrypted password")
	ErrHashPassword              = errors.New("error hashing a password")
	ErrMismatchedHashAndPassword = errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password")
)

func HashPasswordWithGivenCost(password []byte, cost int) (string, error) {
	b, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return "", ErrHashPassword
	}
	return string(b), nil

}

func ComparePasswordWithHashed(password, hashedPassword string) error {

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword)); err != nil {
		return ErrMismatchedHashAndPassword
	}
	return nil

}
