package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/repository"
)

type RefreshTokenUseCase struct {
	UserRepository repository.UserRepository
	JWTSecret      string
}

type RefreshTokenResponse struct {
	AccessToken string
	ExpiresIn   int
}

func NewRefreshTokenUseCase(userRepo repository.UserRepository, jwtSecret string) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		UserRepository: userRepo,
		JWTSecret:      jwtSecret,
	}
}

func (u *RefreshTokenUseCase) Execute(refreshToken string) (*RefreshTokenResponse, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	userID := claims.Subject
	if userID == "" {
		return nil, errors.New("invalid refresh token")
	}

	user, err := u.UserRepository.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	loginUseCase := &LoginUseCase{
		UserRepository: u.UserRepository,
		JWTSecret:      u.JWTSecret,
	}

	accessToken, err := loginUseCase.generateAccessToken(user.Id, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	return &RefreshTokenResponse{
		AccessToken: accessToken,
		ExpiresIn:   900,
	}, nil
}
