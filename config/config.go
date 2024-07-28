package config

import (
	"os"

	"github.com/Rustam2202/message-processor/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel       string
	ServerHostPort string
	PostgresDSN    string
	KafkaBroker    string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal().Caller().Err(err).Msg("Failed to load .env file")
	}

	return &Config{
		LogLevel:       os.Getenv("LOG_LEVEL"),
		ServerHostPort: os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT"),
		PostgresDSN:    os.Getenv("POSTGRES_DSN"),
		KafkaBroker:    os.Getenv("KAFKA_BROKER"),
	}
}
