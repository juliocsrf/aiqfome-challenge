package product

import "github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"

type ProductResponse struct {
	ID        int64   `json:"id"`
	Title     string  `json:"title"`
	Image     string  `json:"image"`
	Price     float64 `json:"price"`
	Rate      float64 `json:"rate"`
	RateCount int64   `json:"rate_count"`
}

type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
	Total    int               `json:"total"`
}

func FromEntity(product *entity.Product) *ProductResponse {
	return &ProductResponse{
		ID:        product.Id,
		Title:     product.Title,
		Image:     product.Image,
		Price:     product.Price,
		Rate:      product.Rate,
		RateCount: product.RateCount,
	}
}

func FromEntities(products []*entity.Product) *ProductListResponse {
	productResponses := make([]ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = *FromEntity(product)
	}

	return &ProductListResponse{
		Products: productResponses,
		Total:    len(productResponses),
	}
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}
