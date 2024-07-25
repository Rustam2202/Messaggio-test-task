package kafka

import (
	"context"
	"log"

	"github.com/Rustam2202/message-processor/models"
	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafkaProducer(broker string) {
	writer = &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    "messages",
		Balancer: &kafka.LeastBytes{},
	}
}

func ProduceMessage(message models.Message) {
	msg := kafka.Message{
		Key:   []byte(string(message.ID)),
		Value: []byte(message.Content),
	}

	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Println("could not write message to kafka:", err)
	}
}
