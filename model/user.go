package model

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Name     string
	Surname  string
	Email    string
	Password string
}

type contextKey string

var (
	ContextUsername = contextKey("username")
)

func (p *User) SetPassword(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.Password = string(hash)
	return nil
}

func (p *User) MatchesPassword(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
func (p *User) CreateJWT(secretKey string) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"email":     p.Email,
		"id":        p.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
