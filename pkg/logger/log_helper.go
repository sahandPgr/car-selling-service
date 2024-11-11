package logger

// logPramsToZerolog is a helper function to convert the extra params to zerolog format
func logPramsToZerolog(extra map[ExtraKey]interface{}) map[string]interface{} {
	params := map[string]interface{}{}
	for k, v := range extra {
		params[string(k)] = v
	}
	return params
}
