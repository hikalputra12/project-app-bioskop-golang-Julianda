package repository

import (
	"app-bioskop/pkg/database"

	"go.uber.org/zap"
)

type Repository struct {
	RegisterRepo RegisterInterface
	log          *zap.Logger
}

func AllRepository(db database.PgxIface, log *zap.Logger) *Repository {
	return &Repository{
		RegisterRepo: NewRegisterRepository(db, log),
	}
}
