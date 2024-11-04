package utils

import (
	"unicode"

	"github.com/sahandPgr/car-selling-service/config"
)

// Define the constants
const (
	lowerCharicters   = "abcdefghijklmnopqrstuvwxyz"
	upperCharicters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharicters = "!@#$%&*"
	numbers           = "0123456789"
	allCharicters     = lowerCharicters + upperCharicters + specialCharicters + numbers
)

// CheckPasswordValid function for check password is valid or not
func CheckPasswordValid(pass string) bool {
	config := config.GetConfig()
	if len(pass) < config.Password.MinLength {
		return false
	}
	if config.Password.IncludeCharacters && !HasLetter(pass) {
		return false
	}
	if config.Password.IncludeNumbers && !HasDigit(pass) {
		return false
	}
	if config.Password.IncludeUppercase && !HasUpper(pass) {
		return false
	}
	if config.Password.IncludeLowercase && !HasLower(pass) {
		return false
	}

	return true
}

// HasDigit function for check password has digit or not
func HasDigit(s string) bool {
	for _, d := range s {
		if unicode.IsDigit(d) {
			return true
		}
	}
	return false
}

// HasLetter function for check password has letter or not
func HasLetter(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

// HasUpper function for check password has upper case or not
func HasUpper(s string) bool {
	for _, c := range s {
		if unicode.IsUpper(c) && unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

// HasLower function for check password has lower case or not
func HasLower(s string) bool {
	for _, c := range s {
		if unicode.IsLower(c) && unicode.IsLetter(c) {
			return true
		}
	}
	return false
}
