package repository

import (
	"app-bioskop/internal/data/entity"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PaymentMethodRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

type PaymentMethodInterface interface {
	GetAllPaymentMethods(ctx context.Context) ([]*entity.PaymentMethod, error)
}

func NewPaymentMethod(db *gorm.DB, log *zap.Logger) PaymentMethodInterface {
	return &PaymentMethodRepo{
		db:  db,
		log: log,
	}
}

// get all payment methods
func (b *PaymentMethodRepo) GetAllPaymentMethods(ctx context.Context) ([]*entity.PaymentMethod, error) {
	var paymentMethod []*entity.PaymentMethod
	result := b.db.WithContext(ctx).Find(&paymentMethod)
	if result.Error != nil {
		return nil, result.Error
	}
	return paymentMethod, nil
}
