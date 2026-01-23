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

type VerifyUsecase struct {
	verifyRepo  repository.VerifyInterface
	Auth        repository.AuthRepoInterface
	SessionRepo repository.SessionRepoInterface
	log         *zap.Logger
}

type VerifyUsecaseInterface interface {
	VerifyOTP(ctx context.Context, req dto.VerifyOTP) (string, error)
}

func NewVerifyUsecase(verifyRepo repository.VerifyInterface, Auth repository.AuthRepoInterface, session repository.SessionRepoInterface, log *zap.Logger) VerifyUsecaseInterface {
	return &VerifyUsecase{
		verifyRepo:  verifyRepo,
		Auth:        Auth,
		SessionRepo: session,
		log:         log,
	}
}

// verify otp
func (b *VerifyUsecase) VerifyOTP(ctx context.Context, req dto.VerifyOTP) (string, error) {
	user, err := b.Auth.FindByEmail(ctx, req.Email)
	if err != nil {
		b.log.Warn("failed: email not found",
			zap.String("email", req.Email))
		return "", errors.New("user not found")
	}
	err = b.verifyRepo.VerifyOTP(ctx, req.Email)
	if err != nil {
		b.log.Error("gagal merubah verify otp", zap.Error(err))
		return "", err
	}
	uuid := utils.NewUUID()
	expiredAt := time.Now()
	lastActive := time.Now()
	fmt.Println("iduser", user.ID)
	session := &entity.Session{
		ExpiredAt:  expiredAt,
		LastActive: lastActive,
		ID:         uuid,
		UserID:     user.ID,
	}
	err = b.SessionRepo.CreateSession(ctx, session)
	if err != nil {
		b.log.Error("failed create session on usercase",
			zap.Error(err),
		)
		return "", errors.New("failed to create session")
	}

	return uuid, nil
}
