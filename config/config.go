package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Conf struct {
	Database Database
	Server   Server
	Auth     Auth
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Schema   string
}

type Server struct {
	Port string
}

type Auth struct {
	JWTSecret string
}

func LoadConfig() (*Conf, error) {
	var err error
	if err = godotenv.Load(); err != nil {
		slog.Warn("No .env file found. Using environment variables")
	}

	conf := &Conf{
		Database: Database{},
		Server:   Server{},
		Auth:     Auth{},
	}

	// conf.Database.ConnString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_SCHEMA"),
	// )

	conf.Database.Host = os.Getenv("DB_HOST")
	conf.Database.Port = os.Getenv("DB_PORT")
	conf.Database.User = os.Getenv("DB_USER")
	conf.Database.Password = os.Getenv("DB_PASSWORD")
	conf.Database.Name = os.Getenv("DB_NAME")
	conf.Database.Schema = os.Getenv("DB_SCHEMA")

	conf.Server.Port = os.Getenv("SERVER_PORT")

	conf.Auth.JWTSecret = os.Getenv("JWT_SECRET")
	if conf.Auth.JWTSecret == "" {
		conf.Auth.JWTSecret = "default-secret-key-change-in-production"
	}

	return conf, nil
}
