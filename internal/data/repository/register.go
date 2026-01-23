package repository

import (
	"app-bioskop/internal/data/entity"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Register struct {
	db  *gorm.DB
	log *zap.Logger
}

type RegisterInterface interface {
	RegisterAccount(ctx context.Context, register *entity.User) error
}

func NewRegisterRepository(db *gorm.DB, log *zap.Logger) RegisterInterface {
	return &Register{
		db:  db,
		log: log,
	}
}

// RegisterAccount registers a new user account in the database.
func (b *Register) RegisterAccount(ctx context.Context, register *entity.User) error {
	result := b.db.Create(&register)
	if result.Error != nil {
		return nil
	}
	return nil
}
