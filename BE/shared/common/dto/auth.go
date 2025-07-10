package dto

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type RegisterResponse struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

type JWTClaims struct {
	UserID   int32  `json:"user_id"`
	Role     string `json:"role"`
	DeviceID string `json:"device_id"`
}
