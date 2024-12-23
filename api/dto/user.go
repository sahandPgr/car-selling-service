package dto

// GetOtpRequest is the payload for requesting an OTP
type GetOtpRequest struct {
	PhoneMobile string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
}

// TokenDetail represents access and refresh token details
type TokenDetail struct {
	AccessTokenExpiredTime  int64  `json:"accessTokenExpiredTime"`
	RefreshTokenExpiredTime int64  `json:"refreshTokenExpiredTime"`
	AccessToken             string `json:"accessToken"`
	RefreshToken            string `json:"refreshToken"`
}

// RegisterUserByUsernameRequest is the payload for user registration by username
type RegisterUserByUsernameRequest struct {
	FirstName string `json:"firstName:" binding:"required,min=3"`
	LastName  string `json:"lastName:" binding:"required,min=6"`
	UserName  string `json:"username" binding:"required,min=5"`
	Email     string `json:"email" binding:"min=6,email"`
	Password  string `json:"password" binding:"required,password,min=6"`
}

// RegisterLoginByMobileRequest is the payload for mobile-based login
type RegisterLoginByMobileRequest struct {
	PhoneMobile string `json:"phoneMobile" binding:"required,mobile,min=11,max=11"`
	Otp         string `json:"otp" binding:"required,min=6,max=6"`
}

// LoginByUsernameRequest is the payload for username-based login
type LoginByUsernameRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=6"`
}
