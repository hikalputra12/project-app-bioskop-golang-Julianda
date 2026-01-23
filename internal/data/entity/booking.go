package entity

// payment ada di booking seat jadi nanti di ganti
type BookingSeat struct {
	Entity
	UserID          int            `gorm:"column:user_id"`
	User            User           `gorm:"foreignKey:UserID;references:ID"`
	ShowTimeId      int            `gorm:"column:showtime_id"`
	ShowTimes       ShowTimes      `gorm:"foreignKey:ShowTimeId;references:ID"`
	SeatId          int            `gorm:"column:seat_id"`
	Seat            Seat           `gorm:"foreignKey:SeatId;references:ID"`
	PaymentMethodID int            `gorm:"column:payment_method_id"`
	PaymentMethod   PaymentMethod  `gorm:"foreignKey:PaymentMethodID;references:ID"`
	Status          string         `gorm:"column:status"`
	PaymentDetails  PaymentDetails `gorm:"-"`
}
type PaymentDetails struct {
	CardNumber string
	CVV        string
	ExpiryDate string
}

type BookingHistory struct {
	MovieTitle string
	Duration   int
	ShowTime   string
	Location   string
}
