package entity

import "time"

type Session struct {
	ID         string
	NewID      string
	UserID     int
	ExpiredAt  time.Time
	RevokedAt  *time.Time
	LastActive time.Time
	CreatedAt  time.Time
}
