package dto

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Phone    string `json:"phone_number" validate:"required,min=10"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
