package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"
)

type ProductRepositoryImpl struct {
	BaseURL string
}

func NewProductRepository() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		BaseURL: "https://fakestoreapi.com",
	}
}

func (p *ProductRepositoryImpl) FindAll() ([]*entity.Product, error) {
	var fakestoreapiResponse []FakestoreapiProductResponse
	var productsResponse []*entity.Product

	url := fmt.Sprintf("%s/products", p.BaseURL)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fakestoreapi error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fakestoreapi error %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&fakestoreapiResponse); err != nil {
		return nil, fmt.Errorf("fakestoreapi error while parsing response: %s", err)
	}

	for _, productResponse := range fakestoreapiResponse {
		productEntity, err := entity.NewProduct(productResponse.ID, productResponse.Title, productResponse.Image, productResponse.Price, productResponse.Rating.Rate, productResponse.Rating.Count)
		if err != nil {
			return nil, fmt.Errorf("fakestoreapi error while creating product entity: %s", err)
		}
		productsResponse = append(productsResponse, productEntity)
	}

	return productsResponse, nil
}

func (p *ProductRepositoryImpl) FindById(id int64) (*entity.Product, error) {
	var fakestoreapiResponse FakestoreapiProductResponse
	var productEntity *entity.Product

	url := fmt.Sprintf("%s/products/%d", p.BaseURL, id)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fakestoreapi error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fakestoreapi error %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&fakestoreapiResponse); err != nil {
		return nil, nil
	}

	if fakestoreapiResponse.ID != 0 {
		productEntity, err = entity.NewProduct(fakestoreapiResponse.ID, fakestoreapiResponse.Title, fakestoreapiResponse.Image, fakestoreapiResponse.Price, fakestoreapiResponse.Rating.Rate, fakestoreapiResponse.Rating.Count)
		if err != nil {
			return nil, fmt.Errorf("fakestoreapi error while creating product entity: %s", err)
		}
	}

	return productEntity, nil
}
