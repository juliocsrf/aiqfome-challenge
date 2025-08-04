package customer

import (
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/repository"
)

type CreateCustomerUseCase struct {
	Repository repository.CustomerRepository
}

func NewCreateCustomerUseCase(repository repository.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		Repository: repository,
	}
}

func (c *CreateCustomerUseCase) Execute(customer *entity.Customer) (*entity.Customer, error) {
	return c.Repository.Create(customer)
}
