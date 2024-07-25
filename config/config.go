package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

type Config struct {
    PostgresDSN string
    KafkaBroker string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    return &Config{
        PostgresDSN: os.Getenv("POSTGRES_DSN"),
        KafkaBroker: os.Getenv("KAFKA_BROKER"),
    }
}