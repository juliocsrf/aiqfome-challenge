package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	authDto "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/dto/auth"
	"github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/utils"
	"github.com/juliocsrf/aiqfome-challenge/internal/usecase/auth"
)

type AuthHandler struct {
	LoginUseCase        *auth.LoginUseCase
	RefreshTokenUseCase *auth.RefreshTokenUseCase
	validator           *validator.Validate
}

func NewAuthHandler(loginUseCase *auth.LoginUseCase, refreshTokenUseCase *auth.RefreshTokenUseCase) *AuthHandler {
	return &AuthHandler{
		LoginUseCase:        loginUseCase,
		RefreshTokenUseCase: refreshTokenUseCase,
		validator:           validator.New(),
	}
}

// Login godoc
// @Summary User login
// @Description Authenticate user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body auth.LoginRequest true "Login credentials"
// @Success 200 {object} auth.LoginResponse
// @Failure 400 {object} auth.ErrorResponse
// @Failure 401 {object} auth.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req authDto.LoginRequest

	_ = json.NewDecoder(r.Body).Decode(&req)

	err := h.validator.Struct(&req)
	if err != nil {
		utils.RespondWithValidationError(w, err)
		return
	}

	result, err := h.LoginUseCase.Execute(req.Email, req.Password)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	response := authDto.LoginResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpiresIn:    result.ExpiresIn,
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Get a new access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body auth.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} auth.RefreshTokenResponse
// @Failure 400 {object} auth.ErrorResponse
// @Failure 401 {object} auth.ErrorResponse
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req authDto.RefreshTokenRequest

	_ = json.NewDecoder(r.Body).Decode(&req)

	err := h.validator.Struct(&req)
	if err != nil {
		utils.RespondWithValidationError(w, err)
		return
	}

	result, err := h.RefreshTokenUseCase.Execute(req.RefreshToken)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	response := authDto.RefreshTokenResponse{
		AccessToken: result.AccessToken,
		ExpiresIn:   result.ExpiresIn,
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}
