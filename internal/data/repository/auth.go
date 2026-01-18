package repository

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/pkg/database"
	"context"

	"go.uber.org/zap"
)

type AuthRepo struct {
	db  database.PgxIface
	log *zap.Logger
}

type AuthRepoInterface interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

func NewAuthRepo(db database.PgxIface, log *zap.Logger) AuthRepoInterface {
	return &AuthRepo{
		db:  db,
		log: log,
	}
}

// FindByEmail retrieves a user by their email address.
func (r *AuthRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	query := `SELECT id, created_at, updated_at, deleted_at, username, email, password_hash, phone_number
        FROM users 
        WHERE email = $1 AND deleted_at IS NULL`
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Phone,
	)
	if err != nil {
		r.log.Error("Database Query Error: failed find user by email",
			zap.Error(err),
			zap.String("query", query),
		)
		return nil, err
	}
	return &user, nil

}
