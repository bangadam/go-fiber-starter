package request

type LoginRequest struct {
	Email    string `json:"email" example:"john.doe@gmail.com" validate:"required,email"`
	Password string `json:"password" example:"12345678" validate:"required,min=8,max=255"`
}

type RegisterRequest struct {
	Email    string `json:"email" example:"john.doe@gmail.com" validate:"required,email"`
	Password string `json:"password" example:"12345678" validate:"required,min=8,max=255"`
}
