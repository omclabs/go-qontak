package auth

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	CreatedAt   int64  `json:"created_at"`
	ExpiredAt   int64  `json:"expired_in"`
}
