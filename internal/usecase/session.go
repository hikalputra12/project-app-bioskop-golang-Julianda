package usecase

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/internal/data/repository"
	"context"

	"go.uber.org/zap"
)

type SessionUsecase struct {
	sessionRepo repository.SessionRepoInterface
	log         *zap.Logger
}

type SessionUsecaseInterface interface {
	CreateSession(ctx context.Context, session *entity.Session) error
	RevokedSession(ctx context.Context, revoke string) error
	ExtendSession(ctx context.Context, session *entity.Session) error
	GetUserIDBySession(ctx context.Context, session string) (int, error)
	IsValid(ctx context.Context, session *entity.Session) (bool, error)
}

func NewSessionUsecase(sessionRepo repository.SessionRepoInterface,
	log *zap.Logger) SessionUsecaseInterface {
	return &SessionUsecase{
		sessionRepo: sessionRepo,
		log:         log,
	}
}

func (s *SessionUsecase) CreateSession(ctx context.Context, session *entity.Session) error {
	if err := s.sessionRepo.CreateSession(ctx, session); err != nil {
		s.log.Error("failed to create session on reposioty", zap.Error(err))
		return err
	}
	return nil
}
func (s *SessionUsecase) RevokedSession(ctx context.Context, revoke string) error {
	if err := s.sessionRepo.RevokedSession(ctx, revoke); err != nil {
		s.log.Error("failed to revoke session on reposioty", zap.Error(err))
		return err
	}
	return nil
}
func (s *SessionUsecase) ExtendSession(ctx context.Context, session *entity.Session) error {
	if err := s.sessionRepo.ExtendSession(ctx, session); err != nil {
		s.log.Error("failed to extend session on reposioty", zap.Error(err))
		return err
	}
	return nil
}
func (s *SessionUsecase) GetUserIDBySession(ctx context.Context, session string) (int, error) {
	userID, err := s.sessionRepo.GetUserIDBySession(ctx, session)
	if err != nil {
		s.log.Error("failed to get user id on reposioty", zap.Error(err))
		return 0, err
	}
	return userID, nil
}

func (s *SessionUsecase) IsValid(ctx context.Context, session *entity.Session) (bool, error) {
	valid, err := s.sessionRepo.IsValid(ctx, session)
	if err != nil {
		s.log.Error("failed to extend session on reposioty", zap.Error(err))
		return false, err
	}
	return valid, nil
}
