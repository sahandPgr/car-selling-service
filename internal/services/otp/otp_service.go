package services

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/constatns"
	"github.com/sahandPgr/car-selling-service/internal/cache"
	serviceerrors "github.com/sahandPgr/car-selling-service/internal/services/service_errors"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
)

// OtpService is a struct for otp service
type OtpService struct {
	log    logger.Logger
	config *config.Config
	redis  *redis.Client
}

// OtpData is a struct for otp data
type OtpData struct {
	Value string
	Used  bool
}

// This function creates a new otp service
func NewOtpService(config *config.Config) *OtpService {
	log := logger.NewLogger(config)
	redisClient := cache.GetRedisClient()
	return &OtpService{
		log:    log,
		config: config,
		redis:  redisClient,
	}
}

// This function sets the otp
func (s *OtpService) SetOtp(otp string, number string) error {
	key := fmt.Sprintf("%s:%s", constatns.RedisOtpDefaultKey, number)
	val := OtpData{
		Value: otp,
		Used:  false,
	}

	res, err := cache.Get[OtpData](s.redis, key)
	if err == nil && !res.Used {
		return serviceerrors.ServiceError{EndUserMessage: serviceerrors.OtpExists}
	} else if err == nil && res.Used {
		return serviceerrors.ServiceError{EndUserMessage: serviceerrors.OtpUsed}
	}

	err = cache.Set(s.redis, key, val, s.config.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}

	return nil
}

// This function validates the otp
func (s *OtpService) ValidatetOtp(otp string, number string) error {
	key := fmt.Sprintf("%s:%s", constatns.RedisOtpDefaultKey, number)
	res, err := cache.Get[OtpData](s.redis, key)

	if err != nil {
		return err
	} else if res.Used {
		return serviceerrors.ServiceError{EndUserMessage: serviceerrors.OtpUsed}
	} else if !res.Used && res.Value != otp {
		return serviceerrors.ServiceError{EndUserMessage: serviceerrors.OtpInvalid}
	} else if !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(s.redis, key, res, s.config.Otp.ExpireTime*time.Second)
		if err != nil {
			return err
		}
	}

	return nil
}
