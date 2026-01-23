package entity

// register ada di user nanti di ubah
type User struct {
	Entity
	Name     string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password_hash"`
	Phone    string `gorm:"column:phone_number"`
}
