package kafka

import (
	"context"
	"encoding/json"

	"github.com/Rustam2202/message-processor/logger"
	"github.com/Rustam2202/message-processor/models"
	"github.com/segmentio/kafka-go"
)

func RunConsumer(broker string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     "messages",
		Partition: 0,
	})

	go func() {
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				logger.Logger.Error().Caller().Err(err).Msg("could not read message")
				continue
			}
			var message models.Message
			if err = json.Unmarshal(msg.Value, &message); err != nil {
				logger.Logger.Error().Caller().Err(err).Msg("could not unmarshal message")
				continue
			}

			tx := models.DB.Model(&models.Message{}).Where("id = ?", message.ID).Update("processed", true)
			if tx.Error != nil {
				logger.Logger.Error().Caller().Err(tx.Error).Msg("could not update message")
				continue
			}

			logger.Logger.Debug().Msgf("Message ID:%d processed: %s", message.ID, message.Content)
			reader.CommitMessages(context.Background(), msg)
		}
	}()
}
