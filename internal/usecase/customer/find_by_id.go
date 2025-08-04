package customer

import (
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/repository"
)

type FindByIdCustomerUseCase struct {
	CustomerRepository  repository.CustomerRepository
	FavoritesRepository repository.FavoritesRepository
	ProductRepository   repository.ProductRepository
}

func NewFindByIdCustomerUseCase(
	customerRepository repository.CustomerRepository,
	favoritesRepository repository.FavoritesRepository,
	productRepository repository.ProductRepository,
) *FindByIdCustomerUseCase {
	return &FindByIdCustomerUseCase{
		CustomerRepository:  customerRepository,
		FavoritesRepository: favoritesRepository,
		ProductRepository:   productRepository,
	}
}

func (f *FindByIdCustomerUseCase) Execute(customerId string) (*entity.Customer, error) {
	customer, err := f.CustomerRepository.FindById(customerId)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, nil
	}

	favoriteProductIds, err := f.FavoritesRepository.FindAllByCustomer(customer)
	if err != nil {
		return customer, nil
	}

	var favoriteProducts []*entity.Product
	for _, productId := range favoriteProductIds {
		if productId != nil {
			product, err := f.ProductRepository.FindById(*productId)
			if err == nil && product != nil {
				favoriteProducts = append(favoriteProducts, product)
			}
		}
	}

	customer.Favorites = favoriteProducts
	return customer, nil
}
