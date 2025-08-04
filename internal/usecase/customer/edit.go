package customer

import (
	"errors"

	"github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/repository"
)

type EditCustomerUseCase struct {
	Repository repository.CustomerRepository
}

func NewEditCustomerUseCase(repository repository.CustomerRepository) *EditCustomerUseCase {
	return &EditCustomerUseCase{
		Repository: repository,
	}
}

func (u *EditCustomerUseCase) Execute(customer *entity.Customer) error {
	customerEntity, err := u.Repository.FindById(customer.Id)
	if err != nil {
		return err
	}

	if customerEntity == nil {
		return errors.New("customer not found")
	}

	customerEntity.Name = customer.Name
	customerEntity.Email = customer.Email

	err = u.Repository.Update(customerEntity)
	if err != nil {
		return err
	}

	return nil
}
