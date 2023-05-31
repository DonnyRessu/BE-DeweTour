package authdto

type RegisterResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
