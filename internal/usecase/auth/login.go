package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	UserRepository repository.UserRepository
	JWTSecret      string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func NewLoginUseCase(userRepo repository.UserRepository, jwtSecret string) *LoginUseCase {
	return &LoginUseCase{
		UserRepository: userRepo,
		JWTSecret:      jwtSecret,
	}
}

func (u *LoginUseCase) Execute(email, password string) (*LoginResponse, error) {
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	accessToken, err := u.generateAccessToken(user.Id, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshToken, err := u.generateRefreshToken(user.Id)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    900,
	}, nil
}

func (u *LoginUseCase) generateAccessToken(userID, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "aiqfome-challenge",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.JWTSecret))
}

func (u *LoginUseCase) generateRefreshToken(userID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "aiqfome-challenge",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.JWTSecret))
}
