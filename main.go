package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Rustam2202/message-processor/config"
	"github.com/Rustam2202/message-processor/handlers"
	"github.com/Rustam2202/message-processor/kafka"
	"github.com/Rustam2202/message-processor/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.LoadConfig()

	if err := models.MustConnectToDatabase(cfg.PostgresDSN); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	kafka.NewKafkaProducer(cfg.KafkaBroker)
	kafka.RunKafkaConsumer(cfg.KafkaBroker)

	server := gin.Default()
	server.POST("/messages", handlers.CreateMessage)
	server.GET("/messages/processed", handlers.GetProcessedMessages)
	server.GET("/messages/count", handlers.GetMessagesCount)
	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
