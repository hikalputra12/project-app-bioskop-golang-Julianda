package entity

import (
	"time"

	"gorm.io/gorm"
)

// nanti di panggil gorm.Entity
type Entity struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//Used for soft deletes (marking records as deleted without actually removing them from the database
}
