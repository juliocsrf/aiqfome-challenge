package entity

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrUserIdEmpty       = errors.New("id cannot be empty")
	ErrUserNameEmpty     = errors.New("name cannot be empty")
	ErrUserEmailEmpty    = errors.New("email cannot be empty")
	ErrUserEmailInvalid  = errors.New("email is invalid")
	ErrUserPasswordEmpty = errors.New("password cannot be empty")
)

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) (*User, error) {
	id := uuid.Must(uuid.NewV7()).String()
	return NewUserWithId(id, name, email, password)
}

func NewUserWithId(id, name, email, password string) (*User, error) {
	var user = &User{
		Id:       id,
		Name:     strings.Join(strings.Fields(name), " "),
		Email:    email,
		Password: password,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.Id == "" {
		return ErrUserIdEmpty
	}

	if u.Name == "" {
		return ErrUserNameEmpty
	}

	if u.Email == "" {
		return ErrUserEmailEmpty
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return ErrUserEmailInvalid
	}

	if u.Password == "" {
		return ErrUserPasswordEmpty
	}

	return nil
}
