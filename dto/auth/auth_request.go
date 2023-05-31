package authdto

type AuthRequest struct {
	FullName string `json:"fullName" form:"fullName"`
	Email    string `json:"email" form:"email" `
	Password string `json:"password" form:"form"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
