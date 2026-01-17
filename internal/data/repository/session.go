package repository

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/pkg/database"
	"context"
	"time"

	"go.uber.org/zap"
)

type SessionRepo struct {
	db  database.PgxIface
	log *zap.Logger
}
type SessionRepoInterface interface {
	CreateSession(ctx context.Context, session *entity.Session) error
	RevokedSession(ctx context.Context, revoke string) error
	ExtendSession(ctx context.Context, session *entity.Session) error
	GetUserIDBySession(ctx context.Context, session string) (int, error)
	IsValid(ctx context.Context, session *entity.Session) (bool, error)
}

func NewSessionRepo(db database.PgxIface, log *zap.Logger) SessionRepoInterface {
	return &SessionRepo{
		db:  db,
		log: log,
	}
}

func (s *SessionRepo) CreateSession(ctx context.Context, session *entity.Session) error {
	query := `INSERT INTO sessions (id,user_id, expired_at ,last_active, created_at) 
			VALUES ($1, $2, $3,$4,$5)`
	now := time.Now()
	expired := time.Now().Add(24 * time.Hour)
	_, err := s.db.Exec(ctx, query, session.ID, session.UserID, expired, now, now)
	if err != nil {
		s.log.Error("failed insert session on database", zap.Error(err), zap.String("query", query))
		return err
	}
	session.ExpiredAt = expired
	session.LastActive = now
	session.CreatedAt = now
	return nil
}

func (s *SessionRepo) RevokedSession(ctx context.Context, revoke string) error {
	query := `UPDATE sessions SET revoked_at=NOW()  WHERE id=$1 AND revoked_at is NULL`

	_, err := s.db.Exec(ctx, query, revoke)
	if err != nil {
		s.log.Error("failed to update revoke on database", zap.Error(err), zap.String("query", query))
		return err
	}
	return nil
}
func (s *SessionRepo) ExtendSession(ctx context.Context, session *entity.Session) error {
	query := `UPDATE sessions SET id=$1, expired_at=$2 ,last_active=NOW() WHERE id=$3 AND revoked_at is NULL`
	expired := time.Now().Add(24 * time.Hour)
	_, err := s.db.Exec(ctx, query, session.NewID, expired, session.ID)
	if err != nil {
		s.log.Error("failed to update revoke on database", zap.Error(err), zap.String("query", query))
		return err
	}
	session.ExpiredAt = expired

	return nil
}

func (s *SessionRepo) IsValid(ctx context.Context, session *entity.Session) (bool, error) {
	query := `SELECT EXISTS(
			  SELECT 1 FROM sessions WHERE id=$1 AND revoked_at is NULL AND expired_at > NOW() )`
	var valid bool
	err := s.db.QueryRow(ctx, query, session.ID).Scan(&valid)
	return valid, err
}

// mendapatkan user id dari session saat itu
func (s *SessionRepo) GetUserIDBySession(ctx context.Context, session string) (int, error) {
	var userID int
	query := `SELECT user_id FROM sessions WHERE id = $1 AND revoked_at IS NULL AND expired_at > NOW()`

	err := s.db.QueryRow(ctx, query, session).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
