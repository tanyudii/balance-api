package config

import (
	"fmt"
	"github.com/tanyudii/balance-api/internal/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB

func GetDatabase() *gorm.DB {
	if db != nil {
		return db
	}

	dbLogLevel := GetConfig().DBLogLevel
	gormConfig := &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormLogLevel(dbLogLevel)),
	}

	dbCon, err := gorm.Open(postgres.Open(getDatabaseDSN()), gormConfig)
	if err != nil {
		logger.Fatalf("failed create connection to database: %v", err)
		return nil
	}

	sqlDB, _ := dbCon.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db = dbCon
	return db
}

func gormLogLevel(level string) gormlogger.LogLevel {
	switch level {
	case "silent":
		return gormlogger.Silent
	case "error":
		return gormlogger.Error
	case "warn":
		return gormlogger.Warn
	default:
		return gormlogger.Info
	}
}

func getDatabaseDSN() string {
	config := GetConfig()
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost, config.DBUsername, config.DBPassword, config.DBDatabase, config.DBPort,
	)
}

func CloseDatabase(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		logger.Fatalf("failed close connection database: %v", err)
		return
	}
	if err = dbSQL.Close(); err != nil {
		logger.Fatalf("failed close connection database: %v", err)
	}
}
