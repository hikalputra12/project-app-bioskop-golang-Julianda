package repository

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/pkg/database"
	"context"
	"time"

	"go.uber.org/zap"
)

type Register struct {
	db  database.PgxIface
	log *zap.Logger
}

type RegisterInterface interface {
	RegisterAccount(ctx context.Context, register *entity.RegisterUser) error
}

func NewRegisterRepository(db database.PgxIface, log *zap.Logger) RegisterInterface {
	return &Register{
		db:  db,
		log: log,
	}
}

// RegisterAccount registers a new user account in the database.
func (r *Register) RegisterAccount(ctx context.Context, register *entity.RegisterUser) error {
	query := `INSERT INTO users (username, email,  phone_number, password_hash, created_at,updated_at) VALUES
($1, $2, $3, $4, $5, $6) RETURNING id`

	now := time.Now()
	err := r.db.QueryRow(ctx, query, register.Name, register.Email, register.Phone, register.Password, now, now).Scan(&register.ID)
	if err != nil {
		r.log.Error("Database Query Error: failed register",
			zap.Error(err),
			zap.String("query", query),
		)
		return err
	}
	register.CreatedAt = now
	register.UpdatedAt = now
	return nil
}
