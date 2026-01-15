package entity

type BookingSeat struct {
	Entity
	UserID        int
	ShowtimeId    int
	SeatId        int
	PaymentMethod int
}
