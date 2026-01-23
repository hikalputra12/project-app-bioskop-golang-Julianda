package entity

type Cinema struct {
	Entity
	Name     string `gorm:"column:name"`
	Location string `gorm:"column:location"`
}
