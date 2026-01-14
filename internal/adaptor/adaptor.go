package adaptor

import (
	"app-bioskop/internal/usecase"

	"go.uber.org/zap"
)

type Adaptor struct {
	RegisterWire  *RegisterAdaptor
	AuthAdaptor   *AuthAdaptor
	CinemaAdaptor *CinemaAdaptor
}

func AllAdaptor(uc usecase.Usecase, log *zap.Logger) *Adaptor {
	return &Adaptor{
		RegisterWire:  NewRegisterAdaptor(uc.RegisterUseCase, log),
		AuthAdaptor:   NewAuthAdaptor(uc.AuthUsecase, uc.SessionUsecase, log),
		CinemaAdaptor: NewCinemaAdaptor(uc.CinemaUsecase, log),
	}
}
