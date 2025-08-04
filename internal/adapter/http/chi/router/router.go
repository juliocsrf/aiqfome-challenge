package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	authHandler "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/handler/auth"
	customerHandler "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/handler/customer"
	favoriteHandler "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/handler/favorite"
	productHandler "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/handler/product"
	appMiddleware "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	CustomerHandler *customerHandler.CustomerHandler
	ProductHandler  *productHandler.ProductHandler
	FavoriteHandler *favoriteHandler.FavoriteHandler
	AuthHandler     *authHandler.AuthHandler
	JWTSecret       string
}

func NewRouter(
	customerHandler *customerHandler.CustomerHandler,
	productHandler *productHandler.ProductHandler,
	favoriteHandler *favoriteHandler.FavoriteHandler,
	authHandler *authHandler.AuthHandler,
	jwtSecret string,
) *Router {
	return &Router{
		CustomerHandler: customerHandler,
		ProductHandler:  productHandler,
		FavoriteHandler: favoriteHandler,
		AuthHandler:     authHandler,
		JWTSecret:       jwtSecret,
	}
}

func (rt *Router) SetupRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(appMiddleware.CORS())
	r.Use(appMiddleware.Logger())
	r.Use(appMiddleware.Recovery())
	r.Use(appMiddleware.RequestID())
	r.Use(appMiddleware.Timeout())
	r.Use(middleware.Compress(5))

	// Swagger documentation
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Route("/api", func(r chi.Router) {
		r.Use(appMiddleware.ContentType())

		// Auth routes (public)
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", rt.AuthHandler.Login)
			r.Post("/refresh", rt.AuthHandler.RefreshToken)
		})

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(appMiddleware.JWTAuth(rt.JWTSecret))

			r.Route("/customers", func(r chi.Router) {
				r.Post("/", rt.CustomerHandler.CreateCustomer)
				r.Get("/{id}", rt.CustomerHandler.GetCustomer)
				r.Put("/{id}", rt.CustomerHandler.UpdateCustomer)
				r.Delete("/{id}", rt.CustomerHandler.DeleteCustomer)

				r.Route("/{customer_id}/favorites", func(r chi.Router) {
					r.Post("/{product_id}", rt.FavoriteHandler.CreateFavorite)
					r.Delete("/{product_id}", rt.FavoriteHandler.DeleteFavorite)
				})
			})

			r.Route("/products", func(r chi.Router) {
				r.Get("/", rt.ProductHandler.GetProducts)
				r.Get("/{id}", rt.ProductHandler.GetProduct)
			})
		})
	})

	return r
}

func (rt *Router) GetAPIRoutes() []RouteInfo {
	return []RouteInfo{
		// Auth routes
		{Method: "POST", Path: "/api/auth/login", Description: "Login user"},
		{Method: "POST", Path: "/api/auth/refresh", Description: "Refresh access token"},

		// Protected routes
		{Method: "POST", Path: "/api/customers", Description: "Create a new customer"},
		{Method: "GET", Path: "/api/customers/{id}", Description: "Get customer by ID"},
		{Method: "PUT", Path: "/api/customers/{id}", Description: "Update customer"},
		{Method: "DELETE", Path: "/api/customers/{id}", Description: "Delete customer"},

		{Method: "GET", Path: "/api/products", Description: "List all products"},
		{Method: "GET", Path: "/api/products/{id}", Description: "Get product by ID"},

		{Method: "POST", Path: "/api/customers/{customer_id}/favorites/{product_id}", Description: "Add product to favorites"},
		{Method: "DELETE", Path: "/api/customers/{customer_id}/favorites/{product_id}", Description: "Remove product from favorites"},
	}
}

type RouteInfo struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	Description string `json:"description"`
}
