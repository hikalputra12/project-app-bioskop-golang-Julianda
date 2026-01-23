package entity

import "time"

type Session struct {
	ID         string     `gorm:"primaryKey;column:id"`
	NewID      string     `gorm:"column:id"`
	UserID     int        `gorm:"column:user_id"`
	User       User       `gorm:"foreignKey:UserID;references:ID"`
	ExpiredAt  time.Time  `gorm:"column:expired_at"`
	RevokedAt  *time.Time `gorm:"column:revoked_at"`
	LastActive time.Time  `gorm:"column:last_active"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
}
