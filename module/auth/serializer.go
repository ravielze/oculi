package auth

type LoginSerializer struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

type RegisterSerializer struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required,ascii,lte=64,gte=8"`
}