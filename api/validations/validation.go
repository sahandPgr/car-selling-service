package validations

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/sahandPgr/car-selling-service/utils"
)

// validationError is a struct for validation errors
type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

// GetValidationErrors is a function for get validation errors
func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, er := range err.(validator.ValidationErrors) {
			var elm ValidationError
			elm.Field = er.Field()
			elm.Tag = er.Tag()
			elm.Value = er.Param()
			elm.Message = er.Error()
			validationErrors = append(validationErrors, elm)
		}
		return &validationErrors
	}
	return nil
}

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
