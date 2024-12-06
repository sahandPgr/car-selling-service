package utils

import (
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// Define the constants
const (
	lowerCharicters   = "abcdefghijklmnopqrstuvwxyz"
	upperCharicters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharicters = "!@#$%&*"
	numbers           = "0123456789"
	allCharicters     = lowerCharicters + upperCharicters + specialCharicters + numbers
)

var (
	cfg           *config.Config = config.GetConfig()
	logg          logger.Logger  = logger.NewLogger(cfg)
	matchFirstCap                = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap                  = regexp.MustCompile("([a-z0-9]([A-Z]))")
)

// CheckPasswordValid function for check password is valid or not
func CheckPasswordValid(pass string) bool {
	if len(pass) < cfg.Password.MinLength {
		return false
	}
	if cfg.Password.IncludeCharacters && !HasLetter(pass) {
		return false
	}
	if cfg.Password.IncludeNumbers && !HasDigit(pass) {
		return false
	}
	if cfg.Password.IncludeUppercase && !HasUpper(pass) {
		return false
	}
	if cfg.Password.IncludeLowercase && !HasLower(pass) {
		return false
	}

	return true
}

// This function for generate password
func GeneratePassword() string {
	var password strings.Builder

	passwordLength := cfg.Password.MinLength + 2
	minSpecialChar := 2
	minNum := 3
	if !cfg.Password.IncludeNumbers {
		minNum = 0
	}

	minUpperCase := 3
	if !cfg.Password.IncludeUppercase {
		minUpperCase = 0
	}

	minLowerCase := 3
	if !cfg.Password.IncludeLowercase {
		minLowerCase = 0
	}

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharicters))
		password.WriteString(string(specialCharicters[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numbers))
		password.WriteString(string(numbers[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharicters))
		password.WriteString(string(upperCharicters[random]))
	}

	//Set lowercase
	for i := 0; i < minLowerCase; i++ {
		random := rand.Intn(len(lowerCharicters))
		password.WriteString(string(lowerCharicters[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharicters))
		password.WriteString(string(allCharicters[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

// this function hash password
func HashedPassword(pass string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		logg.Error(logger.Internal, logger.HashPassword, nil, err.Error())
		return nil, err
	}
	return hash, nil
}

// this function generate otp
func GenerateOtp() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := int(math.Pow(10, float64(cfg.Otp.Digits-1)))
	max := int(math.Pow(10, float64(cfg.Otp.Digits)) - 1)
	res := r.Intn(max-min) + min

	return strconv.Itoa(res)
}

// ToSnakeCase convert models fields to sankecase form.
func ToSnakeCase(str string) (res string) {
	res = matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	res = matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(res)
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
