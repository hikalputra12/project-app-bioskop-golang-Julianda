package adaptor

import (
	"app-bioskop/internal/usecase"

	"go.uber.org/zap"
)

type Adaptor struct {
	RegisterWire *RegisterWire
	log          *zap.Logger
}

func AllAdaptor(uc usecase.Usecase, log *zap.Logger) *Adaptor {
	return &Adaptor{
		RegisterWire: NewRegisterWire(uc.RegisterRepo, log),
		log:          log,
	}
}
