package utils

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func RespondWithValidationError(w http.ResponseWriter, err error) {
	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if !ok {
		RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Validation error"})
		return
	}

	errorMessages := make([]string, len(validationErrors))
	for i, e := range validationErrors {
		errorMessages[i] = getValidationErrorMessage(e)
	}

	RespondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
		"errors": errorMessages,
	})
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func getValidationErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "email":
		return e.Field() + " must be a valid email address"
	default:
		return e.Field() + " is invalid"
	}
}
