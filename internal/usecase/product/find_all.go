package product

import (
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/repository"
)

type FindAllProductUseCase struct {
	Repository repository.ProductRepository
}

func NewFindAllProductUseCase(repository repository.ProductRepository) *FindAllProductUseCase {
	return &FindAllProductUseCase{
		Repository: repository,
	}
}

func (g *FindAllProductUseCase) Execute() ([]*entity.Product, error) {
	products, err := g.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}
