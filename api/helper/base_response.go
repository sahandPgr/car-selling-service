package helper

import "github.com/sahandPgr/car-selling-service/api/validations"

// Define the Base Http response struct
type BaseHttpResponse struct {
	StatusCode int  `json:"statusCode:"`
	Result     any  `json:"result"`
	Error      any  `json:"error"`
	Success    bool `json:"success"`
}

// GetBaseResponse functuon return response without Error and Validation
func GetBaseHttpResponse(result any, success bool, statusCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		StatusCode: statusCode,
	}
}

// GetBaseHttpResponseWithError function return response with Error
func GetBaseHttpResponseWithError(result any, success bool, statusCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		StatusCode: statusCode,
		Error:      err.Error(),
	}
}

// GetBaseHttpResponseWithError function return response with Validation error
func GetBaseHttpResponseWithValidation(result any, success bool, statusCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		StatusCode: statusCode,
		Error:      validations.GetValidationErrors(err),
	}
}
