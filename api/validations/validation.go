package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/sahandPgr/car-selling-service/utils"
)

// IranianPhoneNumberValidator is a custom validator for iranian phone number
func IranianPhoneNumberValidator(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	return utils.IsIranianNumber(val)
}

// PasswordValidator is a custom validator for password
func PasswordValidator(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(string)
	if !ok {
		fl.Param()
		return false
	}
	return utils.CheckPasswordValid(val)
}
