package entity

type BookingSeat struct {
	Entity
	UserID        int
	ShowtimeId    int
	SeatId        int
	PaymentMethod int
	Status        string
}

type PaymentDetails struct {
	CardNumber string
	CVV        string
	ExpiryDate string
}

type Payment struct {
	UserID          int
	BookingId       int
	PaymentMethodID int
	PaymentDetails  PaymentDetails
}

type BookingHistory struct {
	MovieTitle string
	Duration   int
	ShowTime   string
	Location   string
}
