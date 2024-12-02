package serviceerrors

const (
	// OTP
	OtpExists  = "otp exist"
	OtpUsed    = "otp used"
	OtpInvalid = "otp invalid"
	RedisNil   = "redis: nil"

	//JWT
	UnexpectedError = "unexpected error"
	ClaimsNotFound  = "claims not found"

	//User
	EmailExists      = "email exists"
	UsernameExists   = "username exists"
	PermissionDenied = "permission denied"

	//Db
	RecordNotFound = "record not found"
	//Auth
	TokenRequired = "token required"
	TokenExpired  = "token expired"
	TokenInvalid  = "token invalid"
)
