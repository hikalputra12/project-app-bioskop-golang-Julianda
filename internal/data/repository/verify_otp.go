package repository

import (
	"app-bioskop/internal/data/entity"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Verify struct {
	db  *gorm.DB
	log *zap.Logger
}

type VerifyInterface interface {
	VerifyOTP(ctx context.Context, email string) error
}

func NewVerifyRepo(db *gorm.DB, log *zap.Logger) VerifyInterface {
	return &Verify{
		db:  db,
		log: log,
	}
}

// VerifyOTP updates the user's is_verified status to true based on the provided email.
func (b *Verify) VerifyOTP(ctx context.Context, email string) error {
	var user *entity.User
	result := b.db.WithContext(ctx).Where("email=?", email).Model(&user).Update("is_verified", true)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
