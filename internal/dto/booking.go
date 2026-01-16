package dto

type BookingRequest struct {
	Date          string `json:"date" validate:"required,datetime=2006-01-02"`
	Time          string `json:"time" validate:"required,datetime=15:04"`
	SeatIds       []int  `json:"seatIds" validate:"required,min=1"`
	PaymentMethod int    `json:"payment_method,omitempty"`
}

type PaymentDetails struct {
	CardNumber string `json:"card_number" validate:"required,len=16,numeric"`
	CVV        string `json:"cvv" validate:"required,len=3,numeric"`
	ExpiryDate string `json:"expiry_date" validate:"required,datetime=01/06"`
}
type PaymentRequest struct {
	PaymentDetails PaymentDetails `json:"payment_details" validate:"required,dive"`
	PaymentMethod  int            `json:"payment_method,omitempty"`
	BookingId      int            `json:"booking_id" validate:"required"`
}
