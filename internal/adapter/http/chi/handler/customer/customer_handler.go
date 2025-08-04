package customer

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	customerDto "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/dto/customer"
	"github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/utils"
	"github.com/juliocsrf/aiqfome-challenge/internal/usecase/customer"
)

type CustomerHandler struct {
	CreateUseCase   *customer.CreateCustomerUseCase
	FindByIdUseCase *customer.FindByIdCustomerUseCase
	EditUseCase     *customer.EditCustomerUseCase
	DeleteUseCase   *customer.DeleteCustomerUseCase
	validator       *validator.Validate
}

func NewCustomerHandler(
	createUseCase *customer.CreateCustomerUseCase,
	findByIdUseCase *customer.FindByIdCustomerUseCase,
	editUseCase *customer.EditCustomerUseCase,
	deleteUseCase *customer.DeleteCustomerUseCase,
) *CustomerHandler {
	validator := validator.New()
	return &CustomerHandler{
		CreateUseCase:   createUseCase,
		FindByIdUseCase: findByIdUseCase,
		EditUseCase:     editUseCase,
		DeleteUseCase:   deleteUseCase,
		validator:       validator,
	}
}

// CreateCustomer godoc
// @Summary Create a new customer
// @Description Create a new customer with name and email
// @Tags customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body customer.CreateCustomerRequest true "Customer data"
// @Success 201 {object} customer.CustomerResponse
// @Failure 400 {object} customer.ErrorResponse
// @Failure 401 {object} customer.ErrorResponse
// @Router /customers [post]
func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req customerDto.CreateCustomerRequest

	_ = json.NewDecoder(r.Body).Decode(&req)

	err := h.validator.Struct(&req)
	if err != nil {
		utils.RespondWithValidationError(w, err)
		return
	}

	customerEntity, err := req.ToEntity()
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	createdCustomer, err := h.CreateUseCase.Execute(customerEntity)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	response := customerDto.FromEntity(createdCustomer)
	h.writeJSONResponse(w, http.StatusCreated, response)
}

// GetCustomer godoc
// @Summary Get customer by ID
// @Description Get a customer by their ID including favorites
// @Tags customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Customer ID"
// @Success 200 {object} customerDto.CustomerResponse
// @Failure 400 {object} customerDto.ErrorResponse
// @Failure 401 {object} customerDto.ErrorResponse
// @Failure 404 {object} customerDto.ErrorResponse
// @Router /customers/{id} [get]
func (h *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "id")
	if customerID == "" {
		utils.RespondWithValidationError(w, errors.New("customer id is required"))
		return
	}

	customerEntity, err := h.FindByIdUseCase.Execute(customerID)
	if err != nil {
		utils.RespondWithValidationError(w, err)
		return
	}

	if customerEntity == nil {
		h.writeErrorResponse(w, http.StatusNotFound, "customer not found")
		return
	}

	response := customerDto.FromEntity(customerEntity)
	h.writeJSONResponse(w, http.StatusOK, response)
}

// UpdateCustomer godoc
// @Summary Update customer
// @Description Update customer information
// @Tags customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Customer ID"
// @Param request body customer.UpdateCustomerRequest true "Updated customer data"
// @Success 200 {object} customer.CustomerResponse
// @Failure 400 {object} customer.ErrorResponse
// @Failure 401 {object} customer.ErrorResponse
// @Failure 404 {object} customer.ErrorResponse
// @Router /customers/{id} [put]
func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "id")
	if customerID == "" {
		utils.RespondWithValidationError(w, errors.New("customer id is required"))
		return
	}

	var req customerDto.UpdateCustomerRequest

	_ = json.NewDecoder(r.Body).Decode(&req)

	err := h.validator.Struct(&req)
	if err != nil {
		utils.RespondWithValidationError(w, err)
		return
	}

	customerEntity, err := req.ToEntityWithId(customerID)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "invalid request")
		return
	}

	err = h.EditUseCase.Execute(customerEntity)
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}

	response := customerDto.FromEntity(customerEntity)
	h.writeJSONResponse(w, http.StatusOK, response)
}

// DeleteCustomer godoc
// @Summary Delete customer
// @Description Delete a customer by ID
// @Tags customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Customer ID"
// @Success 200 {object} customer.SuccessResponse
// @Failure 400 {object} customer.ErrorResponse
// @Failure 401 {object} customer.ErrorResponse
// @Failure 404 {object} customer.ErrorResponse
// @Router /customers/{id} [delete]
func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "id")
	if customerID == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "customer id is required")
		return
	}

	err := h.DeleteUseCase.Execute(customerID)
	if err != nil {
		if err.Error() == "customer not found" {
			h.writeErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		h.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := customerDto.SuccessResponse{Message: "customer deleted successfully"}
	h.writeJSONResponse(w, http.StatusOK, response)
}

func (h *CustomerHandler) writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *CustomerHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, error string) {
	response := customerDto.ErrorResponse{
		Error: error,
	}
	h.writeJSONResponse(w, statusCode, response)
}
