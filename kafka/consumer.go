package kafka

import (
	"context"
	"log"

	"github.com/Rustam2202/message-processor/config"
	"github.com/Rustam2202/message-processor/models"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func InitKafkaConsumer(broker string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     "messages",
		Partition: 0,
	})

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	cfg := config.LoadConfig()
	if err := models.ConnectDatabase(cfg.PostgresDSN); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	go func() {
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Println("could not read message from kafka:", err)
				continue
			}
			log.Println("received message:", string(msg.Value))
			// Mark message as processed in the database
			tx := models.DB.Update("processed", true)
			if tx.Error != nil {
				log.Println("could not update message in database:", tx.Error)
			}

			reader.CommitMessages(context.Background(), msg)
		}
	}()
}
