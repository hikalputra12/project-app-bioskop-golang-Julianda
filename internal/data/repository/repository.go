package repository

import (
	"app-bioskop/pkg/database"

	"go.uber.org/zap"
)

type Repository struct {
	RegisterRepo RegisterInterface
	AuthRepo     AuthRepoInterface
	SessionRepo  SessionRepoInterface
	CinemaRepo   CinemasRepoInterface
}

func AllRepository(db database.PgxIface, log *zap.Logger) *Repository {
	return &Repository{
		RegisterRepo: NewRegisterRepository(db, log),
		AuthRepo:     NewAuthRepo(db, log),
		SessionRepo:  NewSessionRepo(db, log),
		CinemaRepo:   NewCinemaRepo(db, log),
	}
}
