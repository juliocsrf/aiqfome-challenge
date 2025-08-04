package customer

import (
	"strings"

	"github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"
)

type CreateCustomerRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateCustomerRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (r *CreateCustomerRequest) ToEntity() (*entity.Customer, error) {
	return entity.NewCustomer(
		strings.TrimSpace(r.Name),
		strings.TrimSpace(r.Email),
	)
}

func (r *UpdateCustomerRequest) ToEntityWithId(id string) (*entity.Customer, error) {
	return entity.NewCustomerWithId(
		id,
		strings.TrimSpace(r.Name),
		strings.TrimSpace(r.Email),
	)
}
