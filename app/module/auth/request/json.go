package request

type LoginRequest struct {
	Username string `json:"username" example:"admin1" validate:"required,min=3,max=255"`
	Password string `json:"password" example:"12345678" validate:"required,min=8,max=255"`
}
