package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/juliocsrf/aiqfome-challenge/internal/adapter/database"
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"
	"github.com/lib/pq"
)

type FavoritesRepositoryImpl struct {
	Queries *database.Queries
}

func NewFavoritesRepository(queries *database.Queries) *FavoritesRepositoryImpl {
	return &FavoritesRepositoryImpl{
		Queries: queries,
	}
}

func (f *FavoritesRepositoryImpl) FindAllByCustomer(customer *entity.Customer) ([]*int64, error) {
	var productsIds []*int64
	ctx := context.Background()

	customerUUID, err := uuid.Parse(customer.Id)
	if err != nil {
		return productsIds, nil
	}

	favorites, err := f.Queries.FindAllFavoriteProdutsFromCustomer(ctx, customerUUID)

	if err != nil {
		return productsIds, fmt.Errorf("error while getting customer favorites: %s", err)
	}

	for _, favorite := range favorites {
		productsIds = append(productsIds, &favorite.ProductID)
	}

	return productsIds, nil
}

func (f *FavoritesRepositoryImpl) AddToCustomer(customer *entity.Customer, productId *int64) error {
	ctx := context.Background()
	customerUUID, _ := uuid.Parse(customer.Id)

	err := f.Queries.InsertFavoriteCustomerProduct(ctx, database.InsertFavoriteCustomerProductParams{
		CustomerID: customerUUID,
		ProductID:  *productId,
	})

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return fmt.Errorf("product already in favorites")
			}
		}

		return fmt.Errorf("error while inserting favorite product: %s", err)
	}

	return nil
}

func (f *FavoritesRepositoryImpl) RemoveFromCustomer(customer *entity.Customer, productId *int64) error {
	ctx := context.Background()
	customerUUID, _ := uuid.Parse(customer.Id)

	err := f.Queries.DeleteFavoriteCustomerProduct(ctx, database.DeleteFavoriteCustomerProductParams{
		CustomerID: customerUUID,
		ProductID:  *productId,
	})

	if err != nil {
		return fmt.Errorf("error while deleting favorite product: %s", err)
	}

	return nil
}
