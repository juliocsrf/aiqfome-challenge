package repository

import "github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"

type ProductRepository interface {
	FindAll() ([]*entity.Product, error)
	FindById(id int64) (*entity.Product, error)
}
