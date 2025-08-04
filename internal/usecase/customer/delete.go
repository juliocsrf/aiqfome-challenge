package customer

import (
	"errors"

	"github.com/juliocsrf/aiqfome-challenge/internal/domain/repository"
)

type DeleteCustomerUseCase struct {
	Repository repository.CustomerRepository
}

func NewDeleteCustomerUseCase(repository repository.CustomerRepository) *DeleteCustomerUseCase {
	return &DeleteCustomerUseCase{
		Repository: repository,
	}
}

func (u *DeleteCustomerUseCase) Execute(customerId string) error {
	customer, err := u.Repository.FindById(customerId)
	if err != nil {
		return err
	}

	if customer == nil {
		return errors.New("customer not found")
	}

	err = u.Repository.Delete(customer)
	if err != nil {
		return err
	}

	return nil
}
