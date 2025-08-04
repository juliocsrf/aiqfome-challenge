package repository

import "github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"

type FavoritesRepository interface {
	FindAllByCustomer(*entity.Customer) ([]*int64, error)
	AddToCustomer(*entity.Customer, *int64) error
	RemoveFromCustomer(*entity.Customer, *int64) error
}
