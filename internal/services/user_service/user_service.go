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

// PassChecker is a struct for password checker
type PassChecker struct {
	checker bool   `default:"false"`
	pass    string `default:""`
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
	// Check user's email is exist or not
	if exist := s.existsByEmail(req.Email); exist {
		return &serviceerrors.ServiceError{EndUserMessage: serviceerrors.EmailExists}
	}
	// Check user's username is exist or not
	if exist := s.existsByUsername(req.UserName); exist {
		return &serviceerrors.ServiceError{EndUserMessage: serviceerrors.UsernameExists}
	}

	user := new(models.User)
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Username = req.UserName
	user.Email = req.Email

	passHashed, err := utils.HashedPassword(req.Password)
	if err != nil {
		return err
	}
	user.Password = string(passHashed)
	roleId := s.getDefaultRole()
	tx := s.database.Begin()
	if err = tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&models.UserRole{RoleId: roleId, UserId: user.ID}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

// Register or login a new user by mobile number
func (s *UserService) RegisterLoginByMobile(req *dto.RegisterLoginByMobileRequest) (*dto.TokenDetail, error) {
	if err := s.otp.ValidatetOtp(req.Otp, req.PhoneMobile); err != nil {
		return nil, err
	}

	exist := s.existsByMobileNumber(req.PhoneMobile)
	user := new(models.User)
	user.PhoneMobile = req.PhoneMobile
	user.Username = req.PhoneMobile

	if exist {
		// login
		return s.prepareToken(req.PhoneMobile, PassChecker{checker: false, pass: ""})
	}

	user.Password = string(utils.GeneratePassword())
	roleId := s.getDefaultRole()

	tx := s.database.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Create(&models.UserRole{RoleId: roleId, UserId: user.ID}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return s.prepareToken(req.PhoneMobile, PassChecker{checker: false, pass: ""})
}

// Login by username
func (s *UserService) LoginByUsername(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	return s.prepareToken(req.Username, PassChecker{checker: true, pass: req.Password})
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
func (s *UserService) existsByEmail(email string) bool {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		s.log.Error(logger.Potgres, logger.Select, nil, err.Error())
		return false
	}

	return exists
}

// Check user's username is exist or not
func (s *UserService) existsByUsername(username string) bool {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		s.log.Error(logger.Potgres, logger.Select, nil, err.Error())
		return false
	}

	return exists
}

// Check user's mobile number is exist or not
func (s *UserService) existsByMobileNumber(number string) bool {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("phone_mobile = ?", number).
		Find(&exists).
		Error; err != nil {
		s.log.Error(logger.Potgres, logger.Select, nil, err.Error())
		return false
	}

	return exists
}

// Set Default role for each user
func (s *UserService) getDefaultRole() (roleId int) {
	if err := s.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constatns.DefaultUserRole).
		First(&roleId).
		Error; err != nil {
		s.log.Error(logger.Potgres, logger.Select, nil, err.Error())
		return 0
	}

	return roleId
}

// This function prepares token for user
func (s *UserService) prepareToken(username string, passChecker PassChecker) (token *dto.TokenDetail, err error) {
	u := new(models.User)
	if err = s.database.Model(&models.User{}).
		Where("username = ?", username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).Find(&u).Error; err != nil {
		return nil, err
	}
	if passChecker.checker {
		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passChecker.pass)); err != nil {
			return nil, err
		}
	}
	tokenDto := new(jwt.TokenDto)
	tokenDto.UserrId = u.ID
	tokenDto.Username = u.Username
	tokenDto.FirstName = u.FirstName
	tokenDto.LastName = u.LastName
	tokenDto.Email = u.Email
	tokenDto.PhoneNumber = u.PhoneMobile

	if len(*u.UserRoles) > 0 {
		for _, userRole := range *u.UserRoles {
			tokenDto.Roles = append(tokenDto.Roles, userRole.Role.Name)
		}
	}
	token, err = s.jwt.GenerateToken(tokenDto)
	if err != nil {
		return nil, err
	}
	return token, nil
}
