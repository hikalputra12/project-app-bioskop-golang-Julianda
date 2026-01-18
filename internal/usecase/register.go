package usecase

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/internal/data/repository"
	"app-bioskop/internal/dto"
	"app-bioskop/pkg/utils"
	"context"

	"go.uber.org/zap"
)

type RegisterUseCase struct {
	registerRepo repository.RegisterInterface
	log          *zap.Logger
	emailJob     chan<- utils.EmailJob
}

type RegisterUseCaseInterface interface {
	RegisterAccount(ctx context.Context, req dto.RegisterRequest) error
}

func NewRegisterUseCase(registerRepo repository.RegisterInterface, log *zap.Logger, emailJob chan<- utils.EmailJob) RegisterUseCaseInterface {
	return &RegisterUseCase{
		registerRepo: registerRepo,
		log:          log,
		emailJob:     emailJob,
	}
}

func (u *RegisterUseCase) RegisterAccount(ctx context.Context, req dto.RegisterRequest) error {

	passwordHash := utils.HashPassword(req.Password)
	newUser := &entity.RegisterUser{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: passwordHash,
		Otp:      utils.GenerateRandomNumber(6),
	}
	u.emailJob <- utils.EmailJob{Email: newUser.Email, Otp: newUser.Otp}
	err := u.registerRepo.RegisterAccount(ctx, newUser)
	if err != nil {
		u.log.Error("Usecase Error: failed register account",
			zap.Error(err),
		)
		return err
	}

	return nil
}
