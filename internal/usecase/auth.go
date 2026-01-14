package usecase

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/internal/data/repository"
	"app-bioskop/pkg/utils"
	"context"
	"errors"

	"go.uber.org/zap"
)

type AuthUsecase struct {
	AuthRepo repository.AuthRepoInterface
	log      *zap.Logger
}

type AuthUsecaseInterface interface {
	Login(ctx context.Context, email, password string) (*entity.User, error)
}

func NewAuthUsecase(authRepo repository.AuthRepoInterface, log *zap.Logger) AuthUsecaseInterface {
	return &AuthUsecase{
		AuthRepo: authRepo,
		log:      log,
	}
}

func (u *AuthUsecase) Login(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := u.AuthRepo.FindByEmail(ctx, email)
	if err != nil {
		u.log.Warn("Login attempt failed: email not found",
			zap.String("email", email))
		return nil, errors.New("user not found")
	}
	if !utils.CompareHashAndPassword([]byte(user.Password), []byte(password)) {
		u.log.Warn("Login attempt failed: password is incorrect",
			zap.String("password", password),
		)
		return nil, errors.New("incorrect password")
	}
	return user, nil
}
