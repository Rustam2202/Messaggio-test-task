package kafka

import (
	"context"
	"encoding/json"
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

type messageReq struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

func ProduceMessage(message models.Message) {
	var req messageReq = messageReq{
		ID:      message.ID,
		Content: message.Content,
	}
	messageBytes, err := json.Marshal(req)
	if err != nil {
		log.Println("could not marshal message:", err)
	}

	msg := kafka.Message{
		Value: messageBytes,
	}

	err = writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Println("could not write message to kafka:", err)
	}
}
