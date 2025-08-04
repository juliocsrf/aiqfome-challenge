package wire

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/juliocsrf/aiqfome-challenge/config"
	"github.com/juliocsrf/aiqfome-challenge/internal/adapter/database"
	authHandler "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/handler/auth"
	customerHandler "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/handler/customer"
	favoriteHandler "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/handler/favorite"
	productHandler "github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/handler/product"
	"github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/router"
	productRepo "github.com/juliocsrf/aiqfome-challenge/internal/adapter/repository/fakestoreapi"
	customerRepo "github.com/juliocsrf/aiqfome-challenge/internal/adapter/repository/postgres"
	"github.com/juliocsrf/aiqfome-challenge/internal/domain/repository"
	"github.com/juliocsrf/aiqfome-challenge/internal/usecase/auth"
	"github.com/juliocsrf/aiqfome-challenge/internal/usecase/customer"
	"github.com/juliocsrf/aiqfome-challenge/internal/usecase/favorite"
	"github.com/juliocsrf/aiqfome-challenge/internal/usecase/product"
)

// Database providers
func ProvideQueries(db *sql.DB) *database.Queries {
	return database.New(db)
}

// Repository providers
func ProvideCustomerRepository(queries *database.Queries) repository.CustomerRepository {
	return customerRepo.NewCustomerRepository(queries)
}

func ProvideFavoritesRepository(queries *database.Queries) repository.FavoritesRepository {
	return customerRepo.NewFavoritesRepository(queries)
}

func ProvideUserRepository(queries *database.Queries) repository.UserRepository {
	return customerRepo.NewUserRepository(queries)
}

func ProvideProductRepository() repository.ProductRepository {
	return productRepo.NewProductRepository()
}

// Use case providers
func ProvideCreateCustomerUseCase(repo repository.CustomerRepository) *customer.CreateCustomerUseCase {
	return customer.NewCreateCustomerUseCase(repo)
}

func ProvideFindByIdCustomerUseCase(
	customerRepo repository.CustomerRepository,
	favoritesRepo repository.FavoritesRepository,
	productRepo repository.ProductRepository,
) *customer.FindByIdCustomerUseCase {
	return customer.NewFindByIdCustomerUseCase(customerRepo, favoritesRepo, productRepo)
}

func ProvideEditCustomerUseCase(repo repository.CustomerRepository) *customer.EditCustomerUseCase {
	return customer.NewEditCustomerUseCase(repo)
}

func ProvideDeleteCustomerUseCase(repo repository.CustomerRepository) *customer.DeleteCustomerUseCase {
	return customer.NewDeleteCustomerUseCase(repo)
}

func ProvideFindAllProductUseCase(repo repository.ProductRepository) *product.FindAllProductUseCase {
	return product.NewFindAllProductUseCase(repo)
}

func ProvideFindByIdProductUseCase(repo repository.ProductRepository) *product.FindByIdProductUseCase {
	return product.NewFindByIdProductUseCase(repo)
}

func ProvideCreateFavoriteUseCase(
	favoritesRepo repository.FavoritesRepository,
	customerRepo repository.CustomerRepository,
	productRepo repository.ProductRepository,
) *favorite.CreateFavoriteUseCase {
	return favorite.NewCreateFavoriteUseCase(favoritesRepo, customerRepo, productRepo)
}

func ProvideDeleteFavoriteUseCase(
	favoritesRepo repository.FavoritesRepository,
	customerRepo repository.CustomerRepository,
	productRepo repository.ProductRepository,
) *favorite.DeleteFavoriteUseCase {
	return favorite.NewDeleteFavoriteUseCase(favoritesRepo, customerRepo, productRepo)
}

func ProvideLoginUseCase(userRepo repository.UserRepository, jwtSecret string) *auth.LoginUseCase {
	return auth.NewLoginUseCase(userRepo, jwtSecret)
}

func ProvideRefreshTokenUseCase(userRepo repository.UserRepository, jwtSecret string) *auth.RefreshTokenUseCase {
	return auth.NewRefreshTokenUseCase(userRepo, jwtSecret)
}

// JWT Secret provider
func ProvideJWTSecret(conf *config.Conf) string {
	return conf.Auth.JWTSecret
}

// Handler providers
func ProvideCustomerHandler(
	createUseCase *customer.CreateCustomerUseCase,
	findByIdUseCase *customer.FindByIdCustomerUseCase,
	editUseCase *customer.EditCustomerUseCase,
	deleteUseCase *customer.DeleteCustomerUseCase,
) *customerHandler.CustomerHandler {
	return customerHandler.NewCustomerHandler(createUseCase, findByIdUseCase, editUseCase, deleteUseCase)
}

func ProvideProductHandler(
	findAllUseCase *product.FindAllProductUseCase,
	findByIdUseCase *product.FindByIdProductUseCase,
) *productHandler.ProductHandler {
	return productHandler.NewProductHandler(findAllUseCase, findByIdUseCase)
}

func ProvideFavoriteHandler(
	createUseCase *favorite.CreateFavoriteUseCase,
	deleteUseCase *favorite.DeleteFavoriteUseCase,
) *favoriteHandler.FavoriteHandler {
	return favoriteHandler.NewFavoriteHandler(createUseCase, deleteUseCase)
}

func ProvideAuthHandler(
	loginUseCase *auth.LoginUseCase,
	refreshTokenUseCase *auth.RefreshTokenUseCase,
) *authHandler.AuthHandler {
	return authHandler.NewAuthHandler(loginUseCase, refreshTokenUseCase)
}

// Router provider
func ProvideRouter(
	customerHandler *customerHandler.CustomerHandler,
	productHandler *productHandler.ProductHandler,
	favoriteHandler *favoriteHandler.FavoriteHandler,
	authHandler *authHandler.AuthHandler,
	jwtSecret string,
) *router.Router {
	return router.NewRouter(customerHandler, productHandler, favoriteHandler, authHandler, jwtSecret)
}

// Wire sets
var RepositorySet = wire.NewSet(
	ProvideCustomerRepository,
	ProvideFavoritesRepository,
	ProvideUserRepository,
	ProvideProductRepository,
)

var UseCaseSet = wire.NewSet(
	ProvideCreateCustomerUseCase,
	ProvideFindByIdCustomerUseCase,
	ProvideEditCustomerUseCase,
	ProvideDeleteCustomerUseCase,
	ProvideFindAllProductUseCase,
	ProvideFindByIdProductUseCase,
	ProvideCreateFavoriteUseCase,
	ProvideDeleteFavoriteUseCase,
	ProvideLoginUseCase,
	ProvideRefreshTokenUseCase,
)

var HandlerSet = wire.NewSet(
	ProvideCustomerHandler,
	ProvideProductHandler,
	ProvideFavoriteHandler,
	ProvideAuthHandler,
)

var AllProviders = wire.NewSet(
	ProvideQueries,
	ProvideJWTSecret,
	RepositorySet,
	UseCaseSet,
	HandlerSet,
	ProvideRouter,
)
