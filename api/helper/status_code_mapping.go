package helper

import (
	"net/http"

	serviceerrors "github.com/sahandPgr/car-selling-service/internal/services/service_errors"
)

// Define the StatusCodeMapping
var StatusCodeMapping = map[string]int{

	//OTP
	serviceerrors.OtpExists:  409,
	serviceerrors.OtpUsed:    409,
	serviceerrors.OtpInvalid: 400,
	serviceerrors.RedisNil:   500,

	//JWT
	serviceerrors.UnexpectedError: 500,
	serviceerrors.ClaimsNotFound:  404,

	//User
	serviceerrors.EmailExists:      409,
	serviceerrors.UsernameExists:   409,
	serviceerrors.PermissionDenied: 403,

	//Db
	serviceerrors.RecordNotFound: 404,

	//Auth
	serviceerrors.TokenRequired: 401,
	serviceerrors.TokenExpired:  401,
	serviceerrors.TokenInvalid:  401,
}

// This function convert service error to status code
func ConvertServiceErrorToStatusCode(err error) int {
	val, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return val
}
