package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Rustam2202/message-processor/logger"
)

var DB *gorm.DB

func ConnectToDatabase(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Fatal().Caller().Err(err).Msg("Failed to connect to database")
	}
	err = DB.AutoMigrate(&Message{})
	if err != nil {
		logger.Logger.Fatal().Caller().Err(err).Msg("Failed to auto migrate database")
	}
}
