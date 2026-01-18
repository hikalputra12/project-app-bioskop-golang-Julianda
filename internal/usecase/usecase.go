package usecase

import (
	"app-bioskop/internal/data/repository"
	"app-bioskop/pkg/utils"

	"go.uber.org/zap"
)

type Usecase struct {
	RegisterUseCase      RegisterUseCaseInterface
	AuthUsecase          AuthUsecaseInterface
	SessionUsecase       SessionUsecaseInterface
	CinemaUsecase        CinemaUsecaseInterface
	PaymentMethodUsecase PaymentMethodUsecaseInterface
	BookingUsecase       BookingUsecaseInterface
	emailJob             chan<- utils.EmailJob
	VerifyUsecase        VerifyUsecaseInterface
}

func AllUseCase(repo *repository.Repository, log *zap.Logger, emailJob chan<- utils.EmailJob) Usecase {
	return Usecase{
		RegisterUseCase:      NewRegisterUseCase(repo.RegisterRepo, log, emailJob),
		AuthUsecase:          NewAuthUsecase(repo.AuthRepo, repo.SessionRepo, log),
		SessionUsecase:       NewSessionUsecase(repo.SessionRepo, log),
		CinemaUsecase:        NewCinemaUsecase(repo.CinemaRepo, log),
		PaymentMethodUsecase: NewPaymentMethodUsecase(repo.PaymentMethodRepo, log),
		BookingUsecase:       NewBookingUsecase(repo.BookingRepo, log),
		VerifyUsecase:        NewVerifyUsecase(repo.VerifyRepo, repo.AuthRepo, repo.SessionRepo, log),
	}
}
