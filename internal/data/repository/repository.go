package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	RegisterRepo      RegisterInterface
	AuthRepo          AuthRepoInterface
	SessionRepo       SessionRepoInterface
	CinemaRepo        CinemasRepoInterface
	PaymentMethodRepo PaymentMethodInterface
	BookingRepo       BookingRepoInterface
	VerifyRepo        VerifyInterface
}

func AllRepository(db *gorm.DB, log *zap.Logger) *Repository {
	return &Repository{
		RegisterRepo:      NewRegisterRepository(db, log),
		AuthRepo:          NewAuthRepo(db, log),
		SessionRepo:       NewSessionRepo(db, log),
		CinemaRepo:        NewCinemaRepo(db, log),
		PaymentMethodRepo: NewPaymentMethod(db, log),
		BookingRepo:       NewBookingRepo(db, log),
		VerifyRepo:        NewVerifyRepo(db, log),
	}
}
