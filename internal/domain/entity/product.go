package entity

import (
	"errors"
)

var (
	ErrProductIdInvalid    = errors.New("id must be greater than 0")
	ErrProductTitleEmpty   = errors.New("title cannot be empty")
	ErrProductImageEmpty   = errors.New("image cannot be empty")
	ErrProductPriceInvalid = errors.New("price must be greater than 0")
	ErrProductRateInvalid  = errors.New("rate must be between 0 and 5")
)

type Product struct {
	Id        int64
	Title     string
	Image     string
	Price     float64
	Rate      float64
	RateCount int64
}

func NewProduct(id int64, title, image string, price, rate float64, rateCount int64) (*Product, error) {
	var product = &Product{
		Id:        id,
		Title:     title,
		Image:     image,
		Price:     price,
		Rate:      rate,
		RateCount: rateCount,
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.Id <= 0 {
		return ErrProductIdInvalid
	}

	if p.Title == "" {
		return ErrProductTitleEmpty
	}

	if p.Image == "" {
		return ErrProductImageEmpty
	}

	if p.Price <= 0 {
		return ErrProductPriceInvalid
	}

	if p.Rate < 0 || p.Rate > 5 {
		return ErrProductRateInvalid
	}

	return nil
}
