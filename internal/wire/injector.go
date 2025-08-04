//go:build wireinject

package wire

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/juliocsrf/aiqfome-challenge/config"
	"github.com/juliocsrf/aiqfome-challenge/internal/adapter/http/chi/router"
)

func InitializeApp(db *sql.DB, conf *config.Conf) (*router.Router, error) {
	wire.Build(AllProviders)
	return &router.Router{}, nil
}
