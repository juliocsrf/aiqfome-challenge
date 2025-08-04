// @title AiQFome Challenge API
// @version 1.0
// @description API REST para gerenciamento de clientes, produtos e favoritos
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/juliocsrf/aiqfome-challenge/config"
	"github.com/juliocsrf/aiqfome-challenge/internal/wire"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting application...")

	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	log.Println("Config loaded successfully")

	log.Println("Opening database connection...")
	dbConn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Name,
		conf.Database.Schema,
	))
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer dbConn.Close()

	if err = dbConn.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	log.Println("Database connection opened successfully")

	log.Println("Running database migrations...")
	m, err := migrate.New(
		"file://database/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.Name,
		),
	)
	if err != nil {
		log.Fatalf("Error while running migrations: %v", err)
	}

	if err = m.Up(); err != nil {
		if err.Error() != "no change" {
			log.Fatalf("Error while running migrations: %v", err)
		}
	}
	log.Println("Database migrations completed successfully")

	routerInstance, err := wire.InitializeApp(dbConn, conf)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	handler := routerInstance.SetupRoutes()

	port := ":8080"
	if conf.Server.Port != "" {
		port = ":" + conf.Server.Port
	}

	log.Printf("Starting server on port %s", port)
	log.Println("Available endpoints:")
	for _, route := range routerInstance.GetAPIRoutes() {
		log.Printf("  %s %s - %s", route.Method, route.Path, route.Description)
	}

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
