package services

import (
	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/internal/db"
	services "github.com/sahandPgr/car-selling-service/internal/services/otp"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
	"github.com/sahandPgr/car-selling-service/utils"
	"gorm.io/gorm"
)

// UserService is a struct for user service
type UserService struct {
	log      logger.Logger
	otp      *services.OtpService
	database *gorm.DB
	config   *config.Config
}

// This function creates a new user service
func NewUserService(config *config.Config) *UserService {
	return &UserService{
		log:      logger.NewLogger(config),
		otp:      services.NewOtpService(config),
		database: db.GetDB(),
		config:   config,
	}
}

// This function sends otp
func (s *UserService) SendOtp(dto *dto.GetOtpRequest) error {
	otp := utils.GenerateOtp(s.config)
	err := s.otp.SetOtp(otp, dto.PhoneMobile)
	if err != nil {
		return err
	}

	return nil
}
