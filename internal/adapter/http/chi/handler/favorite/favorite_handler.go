package favorite

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	favoriteDto "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/dto/favorite"
	"github.com/juliocsrf/aiqfome-challenge/internal/usecase/favorite"
)

type FavoriteHandler struct {
	CreateUseCase *favorite.CreateFavoriteUseCase
	DeleteUseCase *favorite.DeleteFavoriteUseCase
}

func NewFavoriteHandler(
	createUseCase *favorite.CreateFavoriteUseCase,
	deleteUseCase *favorite.DeleteFavoriteUseCase,
) *FavoriteHandler {
	return &FavoriteHandler{
		CreateUseCase: createUseCase,
		DeleteUseCase: deleteUseCase,
	}
}

// CreateFavorite godoc
// @Summary Add product to favorites
// @Description Add a product to customer's favorites list
// @Tags favorites
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param customer_id path string true "Customer ID"
// @Param product_id path int true "Product ID"
// @Success 201 {object} favorite.FavoriteResponse
// @Failure 400 {object} favorite.ErrorResponse
// @Failure 401 {object} favorite.ErrorResponse
// @Failure 404 {object} favorite.ErrorResponse
// @Router /customers/{customer_id}/favorites/{product_id} [post]
func (h *FavoriteHandler) CreateFavorite(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "customer_id")
	if customerID == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "customer id is required")
		return
	}

	productIDStr := chi.URLParam(r, "product_id")
	if productIDStr == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "product id is required")
		return
	}

	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "product id must be a number")
		return
	}

	err = h.CreateUseCase.Execute(customerID, productID)
	if err != nil {
		if err.Error() == "customer not found" || err.Error() == "product not found" {
			h.writeErrorResponse(w, http.StatusNotFound, err.Error())
		} else if err.Error() == "product already in favorites" {
			h.writeErrorResponse(w, http.StatusBadRequest, err.Error())
		} else {
			h.writeErrorResponse(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response := favoriteDto.FavoriteResponse{
		CustomerID: customerID,
		ProductID:  productID,
		Message:    "product added to favorites successfully",
	}
	h.writeJSONResponse(w, http.StatusCreated, response)
}

// DeleteFavorite godoc
// @Summary Remove product from favorites
// @Description Remove a product from customer's favorites list
// @Tags favorites
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param customer_id path string true "Customer ID"
// @Param product_id path int true "Product ID"
// @Success 200 {object} favorite.FavoriteResponse
// @Failure 400 {object} favorite.ErrorResponse
// @Failure 401 {object} favorite.ErrorResponse
// @Failure 404 {object} favorite.ErrorResponse
// @Router /customers/{customer_id}/favorites/{product_id} [delete]
func (h *FavoriteHandler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "customer_id")
	if customerID == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "customer id is required")
		return
	}

	productIDStr := chi.URLParam(r, "product_id")
	if productIDStr == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "product id is required")
		return
	}

	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "product id must be a number")
		return
	}

	err = h.DeleteUseCase.Execute(customerID, productID)
	if err != nil {
		if err.Error() == "customer not found" || err.Error() == "product not found" {
			h.writeErrorResponse(w, http.StatusNotFound, err.Error())
		} else {
			h.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response := favoriteDto.FavoriteResponse{
		CustomerID: customerID,
		ProductID:  productID,
		Message:    "product removed from favorites successfully",
	}
	h.writeJSONResponse(w, http.StatusOK, response)
}

func (h *FavoriteHandler) writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *FavoriteHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, error string) {
	response := favoriteDto.ErrorResponse{
		Error: error,
	}
	h.writeJSONResponse(w, statusCode, response)
}
