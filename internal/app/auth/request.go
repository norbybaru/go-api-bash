package auth

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,lte=255"`
	Email    string `json:"email" validate:"required,email,lte=255"`
	Nickname string `json:"nickname" validate:"lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,lte=255"`
}
