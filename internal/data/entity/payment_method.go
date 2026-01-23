package entity

type PaymentMethod struct {
	Entity
	MethodName string `gorm:"column:name"`
	Logo       string `gorm:"column:logo_url"`
}
