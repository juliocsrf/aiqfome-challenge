package favorite

import (
	"errors"

	"github.com/juliocsrf/aiqfome-challenge/internal/domain/repository"
)

type CreateFavoriteUseCase struct {
	FavoritesRepository repository.FavoritesRepository
	CustomerRepository  repository.CustomerRepository
	ProductRepository   repository.ProductRepository
}

func NewCreateFavoriteUseCase(favoritesRepository repository.FavoritesRepository, customerRepository repository.CustomerRepository, productRepository repository.ProductRepository) *CreateFavoriteUseCase {
	return &CreateFavoriteUseCase{
		FavoritesRepository: favoritesRepository,
		CustomerRepository:  customerRepository,
		ProductRepository:   productRepository,
	}
}

func (u *CreateFavoriteUseCase) Execute(customerId string, productId int64) error {
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

	return u.FavoritesRepository.AddToCustomer(customer, &product.Id)
}
