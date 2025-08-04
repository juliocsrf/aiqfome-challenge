package middleware

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func CORS() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})
}

func Logger() func(http.Handler) http.Handler {
	return middleware.Logger
}

func Recovery() func(http.Handler) http.Handler {
	return middleware.Recoverer
}

func ContentType() func(http.Handler) http.Handler {
	return middleware.SetHeader("Content-Type", "application/json")
}

func RequestID() func(http.Handler) http.Handler {
	return middleware.RequestID
}

func Timeout() func(http.Handler) http.Handler {
	return middleware.Timeout(60)
}

func LogInfo(message string) {
	log.Printf("[INFO] %s", message)
}

func LogError(message string, err error) {
	if err != nil {
		log.Printf("[ERROR] %s: %v", message, err)
	} else {
		log.Printf("[ERROR] %s", message)
	}
}
