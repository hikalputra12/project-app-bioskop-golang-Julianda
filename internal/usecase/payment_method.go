package usecase

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/internal/data/repository"
	"context"

	"go.uber.org/zap"
)

type PaymentMethodUsecase struct {
	paymentMethodRepo repository.PaymentMethodInterface
	log               *zap.Logger
}

type PaymentMethodUsecaseInterface interface {
	GetAllPaymentMethods(ctx context.Context) ([]*entity.PaymentMethod, error)
}

// create new payment method usecase
func NewPaymentMethodUsecase(paymentMethodRepo repository.PaymentMethodInterface, log *zap.Logger) PaymentMethodUsecaseInterface {
	return &PaymentMethodUsecase{
		paymentMethodRepo: paymentMethodRepo,
		log:               log,
	}
}

// get all payment methods
func (p *PaymentMethodUsecase) GetAllPaymentMethods(ctx context.Context) ([]*entity.PaymentMethod, error) {
	payment, err := p.paymentMethodRepo.GetAllPaymentMethods(ctx)
	if err != nil {
		p.log.Error("failed to get all payment menthod on repository", zap.Error(err))
		return nil, err
	}
	return payment, nil
}
