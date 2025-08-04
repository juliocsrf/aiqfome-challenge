package repository

import "github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
}
