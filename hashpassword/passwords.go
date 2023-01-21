package hashpassword

import (
	"crypto/rand"
	"crypto/subtle"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
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

func ComparePasswordWithHashed(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil

}

func ComparePasswordWithConfirmPassword(password, ConfirmPassword string) bool {
	return password == ConfirmPassword
}

// TODO plug the hash password below at the point of creating a user and client and service provider
// generate a hashed password using salt
const (
	HashLen = 64
	SaltLen = 32
)

type HashP struct {
	Hash []byte `json:"hash"`
	Salt []byte `json:"salt"`
}

func GenerateSalt() []byte {
	salt := make([]byte, SaltLen)
	_, _ = rand.Read(salt)
	return salt
}

func CreatePassword(password string, salt []byte) ([]byte, error) {
	pass := strings.TrimSpace(password)
	b, err := scrypt.Key([]byte(pass), salt, 13846, 8, 1, HashLen)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func NewPassword(password string) *HashP {
	salt := GenerateSalt()
	Hash, err := CreatePassword(password, salt)
	if err != nil {
		return nil
	}
	return &HashP{
		Hash: Hash,
		Salt: salt,
	}

}

func (h *HashP) Value() (driver.Value, error) {
	j, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}
	return j, err
}

func (h *HashP) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("error unmarshaling json in to struct")
	}
	return json.Unmarshal(b, h)
}

func VerifyPassword(password string, salt, hash []byte) bool {
	newHash, er := CreatePassword(password, salt)
	if er != nil {
		return false
	}
	return subtle.ConstantTimeCompare(newHash, hash) == 1
}

func (h *HashP) IsEqual(pass string) bool {
	b := VerifyPassword(pass, h.Salt, h.Hash)
	return b
}
