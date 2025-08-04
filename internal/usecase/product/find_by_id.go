package product

import (
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/repository"
)

type FindByIdProductUseCase struct {
	Repository repository.ProductRepository
}

func NewFindByIdProductUseCase(repository repository.ProductRepository) *FindByIdProductUseCase {
	return &FindByIdProductUseCase{
		Repository: repository,
	}
}

func (g *FindByIdProductUseCase) Execute(productId int64) (*entity.Product, error) {
	product, err := g.Repository.FindById(productId)
	return product, err
}
