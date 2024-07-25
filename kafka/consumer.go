package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func InitKafkaConsumer(broker string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     "messages",
		Partition: 0,
	})

	go func() {
		for {
			msg, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Println("could not read message from kafka:", err)
				continue
			}

			log.Println("received message:", string(msg.Value))
			// Mark message as processed in the database
		}
	}()
}
