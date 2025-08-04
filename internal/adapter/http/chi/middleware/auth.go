package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/utils"
	"github.com/juliocsrf/aiqfome-challenge/internal/usecase/auth"
)

type contextKey string

const UserIDKey contextKey = "user_id"
const EmailKey contextKey = "email"

func JWTAuth(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "authorization header required"})
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid authorization header format"})
				return
			}

			tokenString := parts[1]

			claims := &auth.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})

			if err != nil || !token.Valid {
				utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid token"})
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, EmailKey, claims.Email)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserIDFromContext(ctx context.Context) string {
	userID, _ := ctx.Value(UserIDKey).(string)
	return userID
}

func GetEmailFromContext(ctx context.Context) string {
	email, _ := ctx.Value(EmailKey).(string)
	return email
}
