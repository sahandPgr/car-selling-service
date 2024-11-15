package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Define the config struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Log      LogConfig
	Cors     CorsConfig
	Password PasswordConfig
}

// Define the Server struct
type ServerConfig struct {
	Port    string
	RunMode string
}

// Define the postgres struct
type PostgresConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	Dbname             string
	SslMode            string
	SetMaxIdleConns    int
	SetMaxOpenConns    int
	SetConnMaxLifetime time.Duration
}

// Define the redis struct
type RedisConfig struct {
	PoolTimeout        time.Duration
	IdleCheckFrequency time.Duration
	WriteTimeout       time.Duration
	ReadTimeout        time.Duration
	DialTimeout        time.Duration
	Db                 int
	PoolSize           int
	Host               string
	Port               string
	Password           string
}

// Define the log struct
type LogConfig struct {
	Level    string
	Encoding string
	Path     string
	Logger   string
}

// Define the password struct
type PasswordConfig struct {
	MinLength         int
	MaxLength         int
	IncludeCharacters bool
	IncludeNumbers    bool
	IncludeUppercase  bool
	IncludeLowercase  bool
}

// Get Cors Config
type CorsConfig struct {
	AllowOrigins string
}

// Get config function
func GetConfig() *Config {
	configPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := loadConfig(configPath, "yml")
	if err != nil {
		log.Fatalf("Error while load config: %v", err)
	}
	config, err := parseConfig(v)
	if err != nil {
		log.Fatalf("Error while parse config: %v", err)
	}
	return config
}

// Get config path
func getConfigPath(env string) string {
	switch env {
	case "docker":
		return "config/config-docker.yml"
	case "production":
		return "config/config-production.yml"
	default:
		return "config/config-development.yml"
	}
}

// Load config with vipet
func loadConfig(filePath string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filePath)
	v.SetConfigType(fileType)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found: %v", err)
			return nil, err
		}
	}
	return v, nil
}

// Parse config to go struct
func parseConfig(v *viper.Viper) (*Config, error) {
	var config Config
	err := v.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to parse config: %v", err)
		return nil, err
	}
	return &config, nil
}
