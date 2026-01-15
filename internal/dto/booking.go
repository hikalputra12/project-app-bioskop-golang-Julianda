package dto

type BookingRequest struct {
	UserID        int    `json:"user_id,omitempty"`
	Date          string `json:"date" validate:"required,datetime=2006-01-02"`
	Time          string `json:"time" validate:"required,datetime=15:04"`
	SeatIds       []int  `json:"seatIds" validate:"required,min=1"`
	PaymentMethod int    `json:"payment_method,omitempty"`
}
