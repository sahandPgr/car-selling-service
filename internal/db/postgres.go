package db

import (
	"fmt"
	"log"
	"time"

	"github.com/sahandPgr/car-selling-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

// This function initialze the Postgres database
func InitialDB(config *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Tehran",
		config.Postgres.Host, config.Postgres.User, config.Postgres.Password, config.Postgres.Dbname, config.Postgres.Port, config.Postgres.SslMode)
	dbClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while connecting to Postgres: %v", err)
	}

	dbSql, _ := dbClient.DB()
	err = dbSql.Ping()

	if err != nil {
		log.Fatalf("Disconnected from Postgres: %v", err)
	}

	dbSql.SetMaxIdleConns(config.Postgres.SetMaxIdleConns)
	dbSql.SetMaxOpenConns(config.Postgres.SetMaxOpenConns)
	dbSql.SetConnMaxLifetime(config.Postgres.SetConnMaxLifetime * time.Minute)
	log.Println("Connected to Postgres :)")
}

// This function returns the DB
func GetDB() *gorm.DB {
	return dbClient
}

// This function closes the DB
func CloseDB() {
	db, _ := dbClient.DB()
	db.Close()
}
