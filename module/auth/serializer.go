package auth

type LoginSerializer struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}

type RegisterSerializer struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}