package favorite

import (
	"errors"

	"github.com/juliocsrf/aiqfome-challenge/internal/domain/repository"
)

type DeleteFavoriteUseCase struct {
	FavoritesRepository repository.FavoritesRepository
	CustomerRepository  repository.CustomerRepository
	ProductRepository   repository.ProductRepository
}

func NewDeleteFavoriteUseCase(favoritesRepository repository.FavoritesRepository, customerRepository repository.CustomerRepository, productRepository repository.ProductRepository) *DeleteFavoriteUseCase {
	return &DeleteFavoriteUseCase{
		FavoritesRepository: favoritesRepository,
		CustomerRepository:  customerRepository,
		ProductRepository:   productRepository,
	}
}

func (u *DeleteFavoriteUseCase) Execute(customerId string, productId int64) error {
	customer, err := u.CustomerRepository.FindById(customerId)
	if err != nil {
		return err
	}

	if customer == nil {
		return errors.New("customer not found")
	}

	product, err := u.ProductRepository.FindById(productId)
	if err != nil {
		return err
	}

	if product == nil {
		return errors.New("product not found")
	}

	return u.FavoritesRepository.RemoveFromCustomer(customer, &product.Id)
}
