package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCustomer_Success(t *testing.T) {
	customer, err := NewCustomer("Júlio Fonseca", "julio.fonseca@gmail.com")

	require.NoError(t, err)
	require.NotNil(t, customer)

	assert.NotEmpty(t, customer.Id)
	assert.Equal(t, "Júlio Fonseca", customer.Name)
	assert.Equal(t, "julio.fonseca@gmail.com", customer.Email)
}

func TestNewCustomerWithId_Success(t *testing.T) {
	customer, err := NewCustomerWithId("123", "Júlio Fonseca", "julio.fonseca@gmail.com")

	require.NoError(t, err)
	require.NotNil(t, customer)

	assert.Equal(t, "123", customer.Id)
	assert.Equal(t, "Júlio Fonseca", customer.Name)
	assert.Equal(t, "julio.fonseca@gmail.com", customer.Email)
}

func TestNewCustomerWithId_EmptyId(t *testing.T) {
	customer, err := NewCustomerWithId("", "Júlio Fonseca", "julio.fonseca@gmail.com")

	assert.Error(t, err)
	assert.Equal(t, ErrCustomerIdEmpty, err)
	assert.Nil(t, customer)
}

func TestNewCustomer_EmptyName(t *testing.T) {
	customer, err := NewCustomer("", "julio.fonseca@gmail.com")

	assert.Error(t, err)
	assert.Equal(t, ErrCustomerNameEmtpy, err)
	assert.Nil(t, customer)
}

func TestNewCustomer_EmptyEmail(t *testing.T) {
	customer, err := NewCustomer("Júlio Fonseca", "")

	assert.Error(t, err)
	assert.Equal(t, ErrCustomerEmailEmpty, err)
	assert.Nil(t, customer)
}

func TestNewCustomer_InvalidEmail(t *testing.T) {
	customer, err := NewCustomer("Júlio Fonseca", "julio.fonseca@@gmail.com")

	assert.Error(t, err)
	assert.Equal(t, ErrCustomerEmailInvalid, err)
	assert.Nil(t, customer)
}
