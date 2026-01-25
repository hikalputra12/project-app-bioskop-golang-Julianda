package repository

import (
	"app-bioskop/internal/data/entity"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

type AuthRepoInterface interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

func NewAuthRepo(db *gorm.DB, log *zap.Logger) AuthRepoInterface {
	return &AuthRepo{
		db:  db,
		log: log,
	}
}

// FindByEmail retrieves a user by their email address.
func (b *AuthRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	result := b.db.WithContext(ctx).Select("id, email, password_hash").Where("email=? AND deleted_at IS NULL", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil

}
