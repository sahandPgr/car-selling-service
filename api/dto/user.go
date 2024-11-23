package dto

type GetOtpRequest struct {
	PhoneMobile string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
}

type TokenDetail struct {
	AccessTokenExpiredTime  int64  `json:"accessTokenExpiredTime"`
	RefreshTokenExpiredTime int64  `json:"refreshTokenExpiredTime"`
	AccessToken             string `json:"accessToken"`
	RefreshToken            string `json:"refreshToken"`
}
