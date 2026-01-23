package entity

import "time"

type Movie struct {
	Entity
	Title       string       `gorm:"column:title"`
	Description string       `gorm:"column:description"`
	Duration    int          `gorm:"column:duration"`
	ReleaseDate time.Time    `gorm:"column:release_date;type:date"`
	Genre       string       `gorm:"column:genre"`
	Director    DirectorInfo `gorm:"column:director_info;type:jsonb;serializer:json"`
	Rating      float32      `gorm:"column:rating"`
	Cast        CastInfo     `gorm:"column:cast_info;type:jsonb;serializer:json"`
}
type DirectorInfo struct {
	Name     string `json:"name"`
	PhotoURL string `json:"photo_url"`
}
type CastInfo struct {
	Name     string `json:"name"`
	PhotoURL string `json:"photo_url"`
}
