package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Rustam2202/message-processor/models"
	"github.com/segmentio/kafka-go"
)

func RunKafkaConsumer(broker string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     "messages",
		Partition: 0,
	})

	go func() {
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Println("could not read message from kafka:", err)
				continue
			}
			var message models.Message
			if err = json.Unmarshal(msg.Value, &message); err != nil {
				log.Println("could not unmarshal message:", err)
				continue
			}

			tx := models.DB.Model(&models.Message{}).Where("id = ?", message.ID).Update("processed", true)
			if tx.Error != nil {
				log.Println("could not update message in database:", tx.Error)
				continue
			}

			reader.CommitMessages(context.Background(), msg)
		}
	}()
}
