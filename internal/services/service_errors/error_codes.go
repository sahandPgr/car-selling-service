package serviceerrors

const (
	// OTP
	OtpExists  = "otp exist"
	OtpUsed    = "otp used"
	OtpInvalid = "otp invalid"

	//JWT
	UnexpectedError = "unexpected error"
	ClaimsNotFound  = "claims not found"
)
