package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/juliocsrf/aiqfome-challenge/internal/adapter/database"
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"
	"github.com/lib/pq"
)

type CustomerRepositoryImpl struct {
	Queries *database.Queries
}

func NewCustomerRepository(queries *database.Queries) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{
		Queries: queries,
	}
}

func (c *CustomerRepositoryImpl) FindById(id string) (*entity.Customer, error) {
	ctx := context.Background()
	customerUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, nil
	}

	customer, err := c.Queries.FindCustomerById(ctx, customerUUID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		return nil, err
	}

	customerEntity, err := entity.NewCustomerWithId(customer.ID.String(), customer.Name, customer.Email)
	if err != nil {
		return nil, fmt.Errorf("error while parsing entity: %s", err)
	}

	return customerEntity, nil
}

func (c *CustomerRepositoryImpl) Create(customer *entity.Customer) (*entity.Customer, error) {
	ctx := context.Background()
	customerUUID, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("error while creating new uuid: %s", err)
	}

	err = c.Queries.InsertCustomer(ctx, database.InsertCustomerParams{
		ID:    customerUUID,
		Name:  customer.Name,
		Email: customer.Email,
	})
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return nil, fmt.Errorf("customer with email %s already exists", customer.Email)
			}
		}

		return nil, fmt.Errorf("error while inserting customer: %s", err)
	}
	customer.Id = customerUUID.String()
	return customer, nil
}

func (c *CustomerRepositoryImpl) Update(customer *entity.Customer) error {
	ctx := context.Background()
	customerUUID, err := uuid.Parse(customer.Id)
	if err != nil {
		return fmt.Errorf("error while parsing customer uuid: %s", err)
	}

	return c.Queries.UpdateCustomer(ctx, database.UpdateCustomerParams{
		Name:  customer.Name,
		Email: customer.Email,
		ID:    customerUUID,
	})
}

func (c *CustomerRepositoryImpl) Delete(customer *entity.Customer) error {
	ctx := context.Background()
	customerUUID, err := uuid.Parse(customer.Id)
	if err != nil {
		uuid.Parse(customer.Id)
	}

	return c.Queries.DeleteCustomer(ctx, customerUUID)
}
