package dtos

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Name      string  `json:"name" binding:"required"`
	Email     string  `json:"email" binding:"required,email"`
	Password  *string `json:"password" binding:"required,excludes= ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=1234567890,containsany=!#$%&'()*+0x2C-./:\"\\;<=>?@[]^_{0x7C}~,min=8,max=128"`
	BirthDate *string `json:"birth_date" binding:"required,datetime=2006-01-02"`
	GenderId  *int64  `json:"gender_id" binding:"required"`
}
