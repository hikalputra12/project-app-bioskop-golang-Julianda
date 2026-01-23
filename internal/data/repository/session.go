package repository

import (
	"app-bioskop/internal/data/entity"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SessionRepo struct {
	db  *gorm.DB
	log *zap.Logger
}
type SessionRepoInterface interface {
	CreateSession(ctx context.Context, session *entity.Session) error
	RevokedSession(ctx context.Context, sessionID string) error
	ExtendSession(ctx context.Context, session *entity.Session) error
	GetUserIDBySession(ctx context.Context, session string) (int, error)
	IsValid(ctx context.Context, session *entity.Session) (bool, error)
}

func NewSessionRepo(db *gorm.DB, log *zap.Logger) SessionRepoInterface {
	return &SessionRepo{
		db:  db,
		log: log,
	}
}

// create new session
func (b *SessionRepo) CreateSession(ctx context.Context, session *entity.Session) error {

	result := b.db.WithContext(ctx).Create(session)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// revoke session
func (b *SessionRepo) RevokedSession(ctx context.Context, sessionID string) error {
	var session *entity.Session
	revoked := time.Now()
	result := b.db.WithContext(ctx).Where("id=? AND revoked_at is NULL", sessionID).Model(&session).Update("revoked_at", revoked)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// extend session
func (b *SessionRepo) ExtendSession(ctx context.Context, session *entity.Session) error {
	result := b.db.WithContext(ctx).Model(&session).Where("id=? AND revoked_at is NULL", session.ID).Updates(map[string]interface{}{"expired_at": session.ExpiredAt, "last_active": session.LastActive, "id": session.NewID})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		fmt.Println("PERINGATAN: Tidak ada session yang diupdate! Cek ID:", session.ID)
		return fmt.Errorf("session not found")
	} else {
		fmt.Println("SUKSES: Session berhasil diperpanjang untuk ID:", session.ID)
	}
	return nil
}

// check is session valid
func (b *SessionRepo) IsValid(ctx context.Context, session *entity.Session) (bool, error) {
	var valid bool
	result := b.db.WithContext(ctx).Raw(`SELECT EXISTS(
			  SELECT 1 FROM sessions WHERE id=$1 AND revoked_at is NULL AND expired_at > NOW() )`, session.ID).Scan(&valid)
	if result.Error != nil {
		return false, result.Error
	}
	return valid, nil
}

// mendapatkan user id dari session saat itu
func (b *SessionRepo) GetUserIDBySession(ctx context.Context, session string) (int, error) {
	var userID int
	result := b.db.WithContext(ctx).Raw(`SELECT user_id FROM sessions WHERE id = $1 AND revoked_at IS NULL AND expired_at > NOW()`, session).Scan(&userID)
	if result.Error != nil {
		return 0, result.Error
	}
	return userID, nil
}
