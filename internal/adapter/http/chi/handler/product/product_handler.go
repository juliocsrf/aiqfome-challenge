package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	productDto "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/dto/product"
	"github.com/juliocsrf/aiqfome-challenge/internal/usecase/product"
)

type ProductHandler struct {
	FindAllUseCase  *product.FindAllProductUseCase
	FindByIdUseCase *product.FindByIdProductUseCase
}

func NewProductHandler(
	findAllUseCase *product.FindAllProductUseCase,
	findByIdUseCase *product.FindByIdProductUseCase,
) *ProductHandler {
	return &ProductHandler{
		FindAllUseCase:  findAllUseCase,
		FindByIdUseCase: findByIdUseCase,
	}
}

// GetProducts godoc
// @Summary List all products
// @Description Get all available products
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} product.ProductListResponse
// @Failure 401 {object} product.ErrorResponse
// @Failure 500 {object} product.ErrorResponse
// @Router /products [get]
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.FindAllUseCase.Execute()
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}

	response := productDto.FromEntities(products)
	h.writeJSONResponse(w, http.StatusOK, response)
}

// GetProduct godoc
// @Summary Get product by ID
// @Description Get a specific product by ID
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} product.ProductResponse
// @Failure 400 {object} product.ErrorResponse
// @Failure 401 {object} product.ErrorResponse
// @Failure 404 {object} product.ErrorResponse
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productIDStr := chi.URLParam(r, "id")
	if productIDStr == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "product id is required")
		return
	}

	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "product id must be a number")
		return
	}

	productEntity, err := h.FindByIdUseCase.Execute(productID)
	if err != nil || productEntity == nil {
		h.writeErrorResponse(w, http.StatusNotFound, "product not found")
		return
	}

	response := productDto.FromEntity(productEntity)
	h.writeJSONResponse(w, http.StatusOK, response)
}

func (h *ProductHandler) writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *ProductHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, error string) {
	response := productDto.ErrorResponse{
		Error: error,
	}
	h.writeJSONResponse(w, statusCode, response)
}
