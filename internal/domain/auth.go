package domain

type SignUpInput struct {
	Username      string `json:"username" binding:"required"`
	Password      string `json:"password" binding:"required,alphanum,min=12,max=48"`
	RegisterToken string `json:"register_token" binding:"required"`
}

type ApplyRegister struct {
	SignUpInput
	IsActive bool
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type JwtIntermediate struct {
	UserID       int    `db:"user_id"`
	PasswordHash string `db:"password_hash"`
	RefreshToken string `db:"refresh_token"`
	IPAddress    string `db:"ip_address"`
	Username     string `db:"username"`
	Role         int    `db:"role"`
}
