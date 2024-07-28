package main

import (
	"github.com/Rustam2202/message-processor/config"
	"github.com/Rustam2202/message-processor/kafka"
	"github.com/Rustam2202/message-processor/logger"
	"github.com/Rustam2202/message-processor/models"
	"github.com/Rustam2202/message-processor/server"
)

func main() {
	cfg := config.LoadConfig()

	logger.NewLogger(cfg.LogLevel)
	models.ConnectToDatabase(cfg.PostgresDSN)

	kafka.NewProducer(cfg.KafkaBroker)
	kafka.RunConsumer(cfg.KafkaBroker)

	server.Run(cfg.ServerHostPort)
}
