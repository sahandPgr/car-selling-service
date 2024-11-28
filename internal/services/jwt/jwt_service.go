package jwt_service

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/constatns"
	serviceerrors "github.com/sahandPgr/car-selling-service/internal/services/service_errors"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
)

// JWTToken struct
type TokenDto struct {
	UserrId     int
	Username    string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Roles       []string
}

type JwtService struct {
	config *config.Config
	log    logger.Logger
}

func NewJwtService(config *config.Config) *JwtService {
	return &JwtService{
		config: config,
		log:    logger.NewLogger(config),
	}
}

// This function generates a new JWT token for the given user.
func (s *JwtService) GenerateToken(token *TokenDto) (*dto.TokenDetail, error) {
	tokenDetail := new(dto.TokenDetail)
	tokenDetail.AccessTokenExpiredTime = time.Now().Add(s.config.Jwt.AccessTokenExpireDuration * time.Minute).Unix()
	tokenDetail.RefreshTokenExpiredTime = time.Now().Add(s.config.Jwt.RefreshTokenExpireDuration * time.Minute).Unix()

	tokenClaims := jwt.MapClaims{
		constatns.UserIdKey:     token.UserrId,
		constatns.UsernameKey:   token.Username,
		constatns.FirstName:     token.FirstName,
		constatns.LastName:      token.LastName,
		constatns.Email:         token.Email,
		constatns.PhoneNumber:   token.PhoneNumber,
		constatns.Roles:         token.Roles,
		constatns.ExpireTimeKey: tokenDetail.AccessTokenExpiredTime,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	var err error
	tokenDetail.AccessToken, err = accessToken.SignedString([]byte(s.config.Jwt.Secret))
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{
		constatns.UserIdKey:     token.UserrId,
		constatns.ExpireTimeKey: tokenDetail.RefreshTokenExpiredTime,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	tokenDetail.RefreshToken, err = refreshToken.SignedString([]byte(s.config.Jwt.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return tokenDetail, nil
}

// This function verifies the given JWT token and returns the claims.
func (s *JwtService) VerifyToken(token string) (*jwt.Token, error) {
	acessToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &serviceerrors.ServiceError{EndUserMessage: serviceerrors.UnexpectedError}
		}
		return []byte(s.config.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	return acessToken, nil
}

// This function returns the claims from the given JWT token.
func (s *JwtService) GetClaims(token string) (map[string]interface{}, error) {
	claimMap := make(map[string]interface{})

	verifiedToken, err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifiedToken.Claims.(jwt.MapClaims)
	if ok && verifiedToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}

	return nil, &serviceerrors.ServiceError{EndUserMessage: serviceerrors.ClaimsNotFound}
}
