package dto

type GetOtpRequest struct {
	PhoneMobile string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
}
