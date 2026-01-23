package usecase

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/internal/data/repository"
	"app-bioskop/internal/dto"
	"app-bioskop/pkg/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type AuthUsecase struct {
	AuthRepo repository.AuthRepoInterface
	Session  repository.SessionRepoInterface
	log      *zap.Logger
}

type AuthUsecaseInterface interface {
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
}

func NewAuthUsecase(authRepo repository.AuthRepoInterface, Session repository.SessionRepoInterface, log *zap.Logger) AuthUsecaseInterface {
	return &AuthUsecase{
		AuthRepo: authRepo,
		Session:  Session,
		log:      log,
	}
}

// Login authenticates a user and creates a session.
func (u *AuthUsecase) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := u.AuthRepo.FindByEmail(ctx, req.Email)
	fmt.Println("iduser", user.ID)
	if err != nil {
		u.log.Warn("Login attempt failed: email not found",
			zap.String("email", req.Email))
		return "", errors.New("user not found")
	}
	if !utils.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) {
		u.log.Warn("Login attempt failed: password is incorrect",
			zap.String("password", req.Password),
		)
		return "", errors.New("incorrect password")
	}
	uuid := utils.NewUUID()

	session := &entity.Session{
		ID:         uuid,
		UserID:     user.ID,
		ExpiredAt:  time.Now().Add(24 * time.Hour),
		LastActive: time.Now(),
	}
	err = u.Session.CreateSession(ctx, session)
	if err != nil {
		u.log.Error("failed create session on usercase",
			zap.Error(err),
		)
		return "", errors.New("failed to create session")
	}

	return uuid, nil
}
