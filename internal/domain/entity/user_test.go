package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser_Success(t *testing.T) {
	user, err := NewUser("Júlio Fonseca", "julio.fonseca@gmail.com", "senha123")

	require.NoError(t, err)
	require.NotNil(t, user)

	assert.NotEmpty(t, user.Id)
	assert.Equal(t, "Júlio Fonseca", user.Name)
	assert.Equal(t, "julio.fonseca@gmail.com", user.Email)
	assert.Equal(t, "senha123", user.Password)
}

func TestNewUserWithId_Success(t *testing.T) {
	user, err := NewUserWithId("123", "Júlio Fonseca", "julio.fonseca@gmail.com", "senha123")

	require.NoError(t, err)
	require.NotNil(t, user)

	assert.Equal(t, "123", user.Id)
	assert.Equal(t, "Júlio Fonseca", user.Name)
	assert.Equal(t, "julio.fonseca@gmail.com", user.Email)
	assert.Equal(t, "senha123", user.Password)
}

func TestNewUserWithId_EmptyId(t *testing.T) {
	user, err := NewUserWithId("", "Júlio Fonseca", "julio.fonseca@gmail.com", "senha123")

	assert.Error(t, err)
	assert.Equal(t, ErrUserIdEmpty, err)
	assert.Nil(t, user)
}

func TestNewUser_EmptyName(t *testing.T) {
	user, err := NewUser("", "julio.fonseca@gmail.com", "senha123")

	assert.Error(t, err)
	assert.Equal(t, ErrUserNameEmpty, err)
	assert.Nil(t, user)
}

func TestNewUser_EmptyEmail(t *testing.T) {
	user, err := NewUser("Júlio Fonseca", "", "senha123")

	assert.Error(t, err)
	assert.Equal(t, ErrUserEmailEmpty, err)
	assert.Nil(t, user)
}

func TestNewUser_InvalidEmail(t *testing.T) {
	user, err := NewUser("Júlio Fonseca", "julio.fonseca@@gmail.com", "senha123")

	assert.Error(t, err)
	assert.Equal(t, ErrUserEmailInvalid, err)
	assert.Nil(t, user)
}

func TestNewUser_EmptyPassword(t *testing.T) {
	user, err := NewUser("Júlio Fonseca", "julio.fonseca@gmail.com", "")

	assert.Error(t, err)
	assert.Equal(t, ErrUserPasswordEmpty, err)
	assert.Nil(t, user)
}
