package usecase

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/internal/data/repository"
	"app-bioskop/pkg/utils"
	"context"

	"go.uber.org/zap"
)

type RegisterUseCase struct {
	registerRepo repository.RegisterInterface
	log          *zap.Logger
}

type RegisterUseCaseInterface interface {
	RegisterAccount(ctx context.Context, register *entity.RegisterUser) error
}

func NewRegisterUseCase(registerRepo repository.RegisterInterface, log *zap.Logger) RegisterUseCaseInterface {
	return &RegisterUseCase{
		registerRepo: registerRepo,
		log:          log,
	}
}

func (u *RegisterUseCase) RegisterAccount(ctx context.Context, register *entity.RegisterUser) error {
	passwordHash := utils.HashPassword(register.Password)
	newUser := &entity.RegisterUser{
		Name:     register.Name,
		Email:    register.Email,
		Phone:    register.Phone,
		Password: passwordHash,
	}
	err := u.registerRepo.RegisterAccount(ctx, newUser)
	if err != nil {
		u.log.Error("Usecase Error: failed register account",
			zap.Error(err),
		)
		return err
	}
	return nil
}
