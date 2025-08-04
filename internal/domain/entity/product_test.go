package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProduct_Success(t *testing.T) {
	product, err := NewProduct(1, "Produto Teste", "https://placehold.co/600x400", 999.99, 4.5, 150)

	require.NoError(t, err)
	require.NotNil(t, product)

	assert.Equal(t, int64(1), product.Id)
	assert.Equal(t, "Produto Teste", product.Title)
	assert.Equal(t, "https://placehold.co/600x400", product.Image)
	assert.Equal(t, 999.99, product.Price)
	assert.Equal(t, 4.5, product.Rate)
	assert.Equal(t, int64(150), product.RateCount)
}

func TestNewProduct_InvalidId(t *testing.T) {
	product, err := NewProduct(0, "Produto Teste", "https://placehold.co/600x400", 999.99, 4.5, 150)

	assert.Error(t, err)
	assert.Equal(t, ErrProductIdInvalid, err)
	assert.Nil(t, product)
}

func TestNewProduct_EmptyTitle(t *testing.T) {
	product, err := NewProduct(1, "", "https://placehold.co/600x400", 999.99, 4.5, 150)

	assert.Error(t, err)
	assert.Equal(t, ErrProductTitleEmpty, err)
	assert.Nil(t, product)
}

func TestNewProduct_EmptyImage(t *testing.T) {
	product, err := NewProduct(1, "Produto Teste", "", 999.99, 4.5, 150)

	assert.Error(t, err)
	assert.Equal(t, ErrProductImageEmpty, err)
	assert.Nil(t, product)
}

func TestNewProduct_InvalidPrice(t *testing.T) {
	product, err := NewProduct(1, "Produto Teste", "https://placehold.co/600x400", 0, 4.5, 150)

	assert.Error(t, err)
	assert.Equal(t, ErrProductPriceInvalid, err)
	assert.Nil(t, product)
}

func TestNewProduct_InvalidRate(t *testing.T) {
	product, err := NewProduct(1, "Produto Teste", "https://placehold.co/600x400", 999.99, 6, 150)

	assert.Error(t, err)
	assert.Equal(t, ErrProductRateInvalid, err)
	assert.Nil(t, product)
}

func TestProduct_Validate_Success(t *testing.T) {
	product := &Product{
		Id:        1,
		Title:     "Produto Teste",
		Image:     "https://example.com/notebook.jpg",
		Price:     2500.00,
		Rate:      4.8,
		RateCount: 200,
	}

	err := product.Validate()
	assert.NoError(t, err)
}

func TestProduct_Validate_InvalidId(t *testing.T) {
	product := &Product{
		Id:        -1,
		Title:     "Produto Teste",
		Image:     "https://example.com/notebook.jpg",
		Price:     2500.00,
		Rate:      4.8,
		RateCount: 200,
	}

	err := product.Validate()
	assert.Error(t, err)
	assert.Equal(t, ErrProductIdInvalid, err)
}

func TestProduct_Validate_RateBoundaries(t *testing.T) {
	product := &Product{
		Id:        1,
		Title:     "Produto Teste",
		Image:     "https://placehold.co/600x400",
		Price:     100.00,
		Rate:      0,
		RateCount: 10,
	}
	err := product.Validate()
	assert.NoError(t, err)

	product.Rate = 5
	err = product.Validate()
	assert.NoError(t, err)

	product.Rate = -0.1
	err = product.Validate()
	assert.Error(t, err)
	assert.Equal(t, ErrProductRateInvalid, err)

	product.Rate = 5.1
	err = product.Validate()
	assert.Error(t, err)
	assert.Equal(t, ErrProductRateInvalid, err)
}
