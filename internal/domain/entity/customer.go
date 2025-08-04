package entity

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrCustomerIdEmpty      = errors.New("id cannot be empty")
	ErrCustomerNameEmtpy    = errors.New("name cannot be empty")
	ErrCustomerEmailEmpty   = errors.New("email cannot be empty")
	ErrCustomerEmailInvalid = errors.New("email is invalid")
)

type Customer struct {
	Id    string
	Name  string
	Email string

	Favorites []*Product
}

func NewCustomer(name, email string) (*Customer, error) {
	id := uuid.Must(uuid.NewV7()).String()
	return NewCustomerWithId(id, name, email)
}

func NewCustomerWithId(id, name, email string) (*Customer, error) {
	var customer = &Customer{
		Id:        id,
		Name:      strings.Join(strings.Fields(name), " "),
		Email:     email,
		Favorites: []*Product{},
	}

	if err := customer.Validate(); err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) Validate() error {
	if c.Id == "" {
		return ErrCustomerIdEmpty
	}

	if c.Name == "" {
		return ErrCustomerNameEmtpy
	}

	if c.Email == "" {
		return ErrCustomerEmailEmpty
	}

	_, err := mail.ParseAddress(c.Email)
	if err != nil {
		return ErrCustomerEmailInvalid
	}

	return nil
}
