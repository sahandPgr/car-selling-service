package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sahandPgr/car-selling-service/api/middlewares"
	"github.com/sahandPgr/car-selling-service/api/routes"
	"github.com/sahandPgr/car-selling-service/api/validations"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/docs"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// This function initializes the server
func InitServer(config *config.Config, log logger.Logger) {
	r := gin.Default()
	registerValidators(log)
	r.Use(middlewares.DefaultMiddlwareLogger(config), middlewares.Cors(config))
	registerRoutes(r, config)
	registerSwagger(r, config)
	err := r.Run(fmt.Sprintf(":%s", config.Server.Port))
	if err != nil {
		log.Fatal(logger.General, logger.Startup, nil, "Failed to start server")
	}
}

// This function registers routes
func registerRoutes(r *gin.Engine, config *config.Config) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		user := v1.Group("/user")
		routes.User(user, config)
	}
}

// This function registers validators
func registerValidators(log logger.Logger) {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianPhoneNumberValidator, true)
		if err != nil {
			log.Error(logger.Validation, logger.Startup, nil, err.Error())
		}
		err = val.RegisterValidation("password", validations.PasswordValidator, true)
		if err != nil {
			log.Error(logger.Validation, logger.Startup, nil, err.Error())
		}

	}
}

func registerSwagger(r *gin.Engine, config *config.Config) {
	docs.SwaggerInfo.Title = "Car Selling Service"
	docs.SwaggerInfo.Description = "Car Selling Service API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", "localhost", config.Server.Port)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
