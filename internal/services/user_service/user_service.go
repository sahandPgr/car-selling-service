package services

import (
	"github.com/sahandPgr/car-selling-service/api/dto"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/constatns"
	"github.com/sahandPgr/car-selling-service/internal/db"
	"github.com/sahandPgr/car-selling-service/internal/models"
	"github.com/sahandPgr/car-selling-service/internal/services/jwt"
	services "github.com/sahandPgr/car-selling-service/internal/services/otp"
	serviceerrors "github.com/sahandPgr/car-selling-service/internal/services/service_errors"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
	"github.com/sahandPgr/car-selling-service/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService is a struct for user service
type UserService struct {
	log      logger.Logger
	otp      *services.OtpService
	database *gorm.DB
	config   *config.Config
	jwt      *jwt.JwtService
}

// This function creates a new user service
func NewUserService(config *config.Config) *UserService {
	return &UserService{
		log:      logger.NewLogger(config),
		otp:      services.NewOtpService(config),
		database: db.GetDB(),
		config:   config,
		jwt:      jwt.NewJwtService(config),
	}
}

// Register a new user
func (s *UserService) RegisterByUsername(req *dto.RegisterUserByUsernameRequest) error {
	exist, err := s.existsByEmail(req.Email)
	if err != nil {
		return err
	}
	if exist {
		return &serviceerrors.ServiceError{EndUserMessage: serviceerrors.EmailExists}
	}
	exist, err = s.existsByUsername(req.UserName)
	if err != nil {
		return err
	}
	if exist {
		return &serviceerrors.ServiceError{EndUserMessage: serviceerrors.UsernameExists}
	}

	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.UserName,
		Email:     req.Email,
	}
	passHashed, err := utils.HashedPassword(req.Password)
	if err != nil {
		return err
	}
	user.Password = string(passHashed)
	roleId, err := s.getDefaultRole()
	if err != nil {
		return err
	}
	tx := s.database.Begin()
	err = tx.Create(user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: user.ID}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *UserService) RegisterLoginByMobile(req *dto.RegisterLoginByMobileRequest) (*dto.TokenDetail, error) {
	err := s.otp.ValidatetOtp(req.Otp, req.PhoneMobile)
	if err != nil {
		return nil, err
	}
	exist, err := s.existsByMobileNumber(req.PhoneMobile)
	user := models.User{
		PhoneMobile: req.PhoneMobile, Username: req.PhoneMobile,
	}
	if exist {
		// login
		var u models.User
		err = s.database.Model(&models.User{}).
			Where("username = ?", user.Username).
			Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
				return tx.Preload("Role")
			}).
			Find(&u).Error
		if err != nil {
			return nil, err
		}
		tokenDto := jwt.TokenDto{
			UserrId:     u.ID,
			Username:    u.Username,
			FirstName:   u.FirstName,
			LastName:    u.LastName,
			Email:       u.Email,
			PhoneNumber: u.PhoneMobile,
		}
		if len(*u.UserRoles) > 0 {
			for _, userRole := range *u.UserRoles {
				tokenDto.Roles = append(tokenDto.Roles, userRole.Role.Name)
			}
		}
		token, err := s.jwt.GenerateToken(&tokenDto)
		if err != nil {
			return nil, err
		}
		return token, nil

	}

	user.Password = string(utils.GeneratePassword())
	roleId, err := s.getDefaultRole()
	if err != nil {
		return nil, err
	}
	tx := s.database.Begin()
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: user.ID}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	var us models.User
	err = s.database.Model(&models.User{}).
		Where("username = ?", user.Username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&us).Error
	if err != nil {
		return nil, err
	}
	tokenDto := jwt.TokenDto{
		UserrId:     user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneMobile,
	}
	if len(*us.UserRoles) > 0 {
		for _, userRole := range *us.UserRoles {
			tokenDto.Roles = append(tokenDto.Roles, userRole.Role.Name)
		}
	}
	token, err := s.jwt.GenerateToken(&tokenDto)
	if err != nil {
		return nil, err
	}

	return token, nil

}

func (s *UserService) LoginByUsername(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	var user models.User
	err := s.database.Model(&models.User{}).
		Where("username = ?", req.Username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	tokenDto := jwt.TokenDto{
		UserrId:     user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneMobile,
	}
	if len(*user.UserRoles) > 0 {
		for _, userRole := range *user.UserRoles {
			tokenDto.Roles = append(tokenDto.Roles, userRole.Role.Name)
		}
	}
	token, err := s.jwt.GenerateToken(&tokenDto)
	if err != nil {
		return nil, err
	}

	return token, nil

}

// This function sends otp
func (s *UserService) SendOtp(dto *dto.GetOtpRequest) error {
	otp := utils.GenerateOtp()
	err := s.otp.SetOtp(otp, dto.PhoneMobile)
	if err != nil {
		return err
	}

	return nil
}

// Check user's email is exist or not
func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		s.log.Error(logger.Potgres, logger.Select, nil, err.Error())
		return false, err
	}

	return exists, nil
}

// Check user's username is exist or not
func (s *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		s.log.Error(logger.Potgres, logger.Select, nil, err.Error())
		return false, err
	}

	return exists, nil
}

// Check user's mobile number is exist or not
func (s *UserService) existsByMobileNumber(number string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("phone_mobile = ?", number).
		Find(&exists).
		Error; err != nil {
		s.log.Error(logger.Potgres, logger.Select, nil, err.Error())
		return false, err
	}

	return exists, nil
}

// Set Default role for each user
func (s *UserService) getDefaultRole() (roleId int, err error) {
	if err = s.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constatns.DefaultUserRole).
		First(&roleId).
		Error; err != nil {
		s.log.Error(logger.Potgres, logger.Select, nil, err.Error())
		return 0, err
	}

	return roleId, nil
}
