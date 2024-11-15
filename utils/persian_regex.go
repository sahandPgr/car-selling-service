package utils

import (
	"log"
	"regexp"
)

// Define the constants
const (
	iranianNumberPattern = `^09(1[0-9]|2[0-9]|3[0-9]|9[0-9])[0-9]{7}$`
)

// IsIranianNumber function for check phone number is iranian or not
func IsIranianNumber(number string) bool {
	res, err := regexp.MatchString(iranianNumberPattern, number)
	if err != nil {
		log.Printf("Error while phone number regex pattern: %s", err.Error())
	}
	return res
}
