package usecase

import (
	"app-bioskop/internal/data/repository"

	"go.uber.org/zap"
)

type Usecase struct {
	RegisterUseCase      RegisterUseCaseInterface
	AuthUsecase          AuthUsecaseInterface
	SessionUsecase       SessionUsecaseInterface
	CinemaUsecase        CinemaUsecaseInterface
	PaymentMethodUsecase PaymentMethodUsecaseInterface
	BookingUsecase       BookingUsecaseInterface
}

func AllUseCase(repo *repository.Repository, log *zap.Logger) Usecase {
	return Usecase{
		RegisterUseCase:      NewRegisterUseCase(repo.RegisterRepo, log),
		AuthUsecase:          NewAuthUsecase(repo.AuthRepo, repo.SessionRepo, log),
		SessionUsecase:       NewSessionUsecase(repo.SessionRepo, log),
		CinemaUsecase:        NewCinemaUsecase(repo.CinemaRepo, log),
		PaymentMethodUsecase: NewPaymentMethodUsecase(repo.PaymentMethodRepo, log),
		BookingUsecase:       NewBookingUsecase(repo.BookingRepo, log),
	}
}
