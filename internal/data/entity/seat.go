package entity

import "time"

type Seat struct {
	Entity
	SeatNumber string `gorm:"column:seat_number"`
	CinemaID   int    `gorm:"column:cinema_id"`
	Cinema     Cinema `gorm:"foreignKey:CinemaID;references:ID"`
	IsAvaiable bool   `gorm:"column:is_available"`
}
type ShowTimes struct {
	Entity
	MovieID   int       `gorm:"column:movie_id"`
	Movie     Movie     `gorm:"foreignKey:MovieID;references:ID"`
	CinemaID  int       `gorm:"column:cinema_id"`
	Cinema    Cinema    `gorm:"foreignKey:CinemaID;references:ID"`
	StartTime time.Time `gorm:"column:start_time"`
	EndtTime  time.Time `gorm:"column:end_time"`
	price     float64   `gorm:"column:price"`
}
