package usecase

import (
	"app-bioskop/internal/data/repository"

	"go.uber.org/zap"
)

type Usecase struct {
	RegisterRepo repository.RegisterInterface
	log          *zap.Logger
}

func AllUseCase(repo repository.Repository, log *zap.Logger) Usecase {
	return Usecase{
		RegisterRepo: NewRegisterUseCase(repo.RegisterRepo, log),
		log:          log,
	}
}
