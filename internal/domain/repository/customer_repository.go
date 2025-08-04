package repository

import "github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"

type CustomerRepository interface {
	FindById(id string) (*entity.Customer, error)
	Create(*entity.Customer) (*entity.Customer, error)
	Update(*entity.Customer) error
	Delete(*entity.Customer) error
}
